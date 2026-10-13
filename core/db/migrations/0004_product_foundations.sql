-- SPDX-License-Identifier: AGPL-3.0-or-later

ALTER TABLE documents ADD COLUMN title_owner TEXT NOT NULL DEFAULT 'auto'
    CHECK (title_owner IN ('auto', 'rule', 'human'));

ALTER TABLE share_links ADD COLUMN visibility TEXT NOT NULL DEFAULT 'network'
    CHECK (visibility IN ('network', 'public'));
ALTER TABLE users ADD COLUMN locale TEXT NOT NULL DEFAULT '';

ALTER TABLE custom_fields ADD COLUMN managed_by TEXT NOT NULL DEFAULT ''
    CHECK (managed_by IN ('', 'zugferd'));

-- Fields created by e-invoice extraction are part of that pipeline's storage
-- contract. Adopt matching rows from existing archives without claiming a
-- conflicting user-defined field whose type would break extraction.
UPDATE custom_fields
SET managed_by = 'zugferd'
WHERE (name = 'invoice_number'      AND data_type = 'text')
   OR (name = 'invoice_date'        AND data_type = 'date')
   OR (name = 'invoice_currency'    AND data_type = 'text')
   OR (name = 'invoice_total_gross' AND data_type = 'monetary')
   OR (name = 'invoice_total_net'   AND data_type = 'monetary')
   OR (name = 'invoice_seller'      AND data_type = 'text')
   OR (name = 'invoice_seller_vat'  AND data_type = 'text')
   OR (name = 'invoice_buyer'       AND data_type = 'text');

CREATE TRIGGER custom_fields_managed_immutable
BEFORE UPDATE OF name, data_type, extra_data, managed_by ON custom_fields
WHEN OLD.managed_by <> ''
 AND (NEW.name IS NOT OLD.name
      OR NEW.data_type IS NOT OLD.data_type
      OR NEW.extra_data IS NOT OLD.extra_data
      OR NEW.managed_by IS NOT OLD.managed_by)
BEGIN SELECT RAISE(ABORT, 'managed custom field is immutable'); END;

CREATE TRIGGER custom_fields_managed_delete
BEFORE DELETE ON custom_fields
WHEN OLD.managed_by <> ''
BEGIN SELECT RAISE(ABORT, 'managed custom field cannot be deleted'); END;

CREATE TRIGGER share_links_visibility_immutable
BEFORE UPDATE OF visibility ON share_links
WHEN NEW.visibility IS NOT OLD.visibility
BEGIN SELECT RAISE(ABORT, 'share visibility is immutable'); END;

CREATE TRIGGER share_links_visibility_replace
BEFORE INSERT ON share_links
WHEN EXISTS (
    SELECT 1 FROM share_links
    WHERE id = NEW.id AND visibility IS NOT NEW.visibility
)
BEGIN SELECT RAISE(ABORT, 'share visibility is immutable'); END;

-- Existing explicit titles are human-owned unless they still match the first
-- acquisition's generated filename title. Imported Paperless titles remain
-- human-owned because their provenance cannot distinguish curated values.
UPDATE documents SET title_owner = 'human';
UPDATE documents
SET title_owner = 'auto'
WHERE title = 'Untitled'
  AND NOT EXISTS (
      SELECT 1 FROM document_sources
      WHERE document_id = documents.id AND kind = 'import'
  );

WITH RECURSIVE
first_sources(document_id, kind, detail) AS (
    SELECT s.document_id, s.kind, replace(s.detail, '\', '/')
    FROM document_sources s
    WHERE s.id = (
        SELECT first.id
        FROM document_sources first
        WHERE first.document_id = s.document_id
        ORDER BY first.observed_at, first.id
        LIMIT 1
    )
      AND s.kind <> 'import'
),
basenames(document_id, kind, name) AS (
    SELECT document_id, kind, detail FROM first_sources
    UNION ALL
    SELECT document_id, kind, substr(name, instr(name, '/') + 1)
    FROM basenames
    WHERE instr(name, '/') > 0
),
source_names(document_id, kind, name) AS (
    SELECT document_id, kind, name
    FROM basenames
    WHERE instr(name, '/') = 0 AND name <> ''
),
reversed_names(document_id, kind, name, reversed, position) AS (
    SELECT document_id, kind, name, '', length(name)
    FROM source_names
    UNION ALL
    SELECT document_id, kind, name,
           reversed || substr(name, position, 1), position - 1
    FROM reversed_names
    WHERE position > 0
),
derived_titles(document_id, title) AS (
    SELECT document_id,
           CASE
               WHEN kind = 'watched_folder' AND instr(reversed, '.') > 0 THEN
                   CASE
                       WHEN length(name) = instr(reversed, '.') THEN 'Untitled'
                       ELSE substr(name, 1, length(name) - instr(reversed, '.'))
                   END
               WHEN lower(name) LIKE '%.jpeg' THEN substr(name, 1, length(name) - 5)
               WHEN lower(name) LIKE '%.tiff' THEN substr(name, 1, length(name) - 5)
               WHEN lower(name) LIKE '%.webp' THEN substr(name, 1, length(name) - 5)
               WHEN lower(name) LIKE '%.pdf'
                 OR lower(name) LIKE '%.png'
                 OR lower(name) LIKE '%.jpg'
                 OR lower(name) LIKE '%.tif'
                   THEN substr(name, 1, length(name) - 4)
               ELSE name
           END
    FROM reversed_names
    WHERE position = 0
)
UPDATE documents
SET title_owner = 'auto'
WHERE EXISTS (
    SELECT 1 FROM derived_titles
    WHERE derived_titles.document_id = documents.id
      AND derived_titles.title = documents.title
);

DROP TRIGGER documents_title_revision;
CREATE TRIGGER documents_title_revision AFTER UPDATE OF title, title_owner ON documents
BEGIN UPDATE documents SET title_revision = title_revision + 1 WHERE id = NEW.id; END;

-- Watched folders are read-only sources. This checkpoint belongs to Suchi and
-- prevents unchanged source files from being rehashed on every reconciliation.
CREATE TABLE fs_watch_observations (
    root_path         TEXT NOT NULL,
    owner_id          INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    system_id         INTEGER NOT NULL REFERENCES jd_systems(id) ON DELETE CASCADE,
    relative_path     TEXT NOT NULL,
    file_size         INTEGER NOT NULL CHECK (file_size >= 0),
    file_mtime_ns     INTEGER NOT NULL,
    sidecar_size      INTEGER NOT NULL DEFAULT -1 CHECK (sidecar_size >= -1),
    sidecar_mtime_ns  INTEGER NOT NULL DEFAULT 0,
    state             TEXT NOT NULL CHECK (state IN ('ingested', 'ignored', 'failed')),
    document_id       INTEGER REFERENCES documents(id) ON DELETE SET NULL,
    last_error        TEXT NOT NULL DEFAULT '',
    updated_at        INTEGER NOT NULL,
    PRIMARY KEY (root_path, owner_id, system_id, relative_path)
) STRICT;

-- Presets no longer seed automations. Editing a built-in rule already forked a
-- user-owned copy, so the remaining preset rows carry no user edits.
DELETE FROM automation_asks WHERE automation_id IN
    (SELECT id FROM automations WHERE COALESCE(preset_slug, '') <> '');
DELETE FROM automation_actions WHERE automation_id IN
    (SELECT id FROM automations WHERE COALESCE(preset_slug, '') <> '');
DELETE FROM automation_triggers WHERE automation_id IN
    (SELECT id FROM automations WHERE COALESCE(preset_slug, '') <> '');
DELETE FROM automations WHERE COALESCE(preset_slug, '') <> '';
DROP INDEX idx_automations_preset_slug;
ALTER TABLE automations DROP COLUMN preset_slug;
