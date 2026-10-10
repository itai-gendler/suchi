-- SPDX-License-Identifier: AGPL-3.0-or-later

ALTER TABLE documents ADD COLUMN title_owner TEXT NOT NULL DEFAULT 'auto'
    CHECK (title_owner IN ('auto', 'rule', 'human'));

ALTER TABLE share_links ADD COLUMN visibility TEXT NOT NULL DEFAULT 'network'
    CHECK (visibility IN ('network', 'public'));

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
