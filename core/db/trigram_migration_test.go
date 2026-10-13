// SPDX-License-Identifier: AGPL-3.0-or-later

package db_test

import (
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/db"
	"github.com/johnnybravo-xyz/suchi/core/db/migrations"
)

// Documents stored before the trigram index existed are indexed by the
// upgrade itself, so substring search works without a manual rebuild.
func TestUpgradeIndexesExistingDocumentsForSubstringSearch(t *testing.T) {
	d, log := openSetDB(t)
	stable, err := db.LoadMigrations(migrations.FS, ".")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Migrate(t.Context(), d, stable[:3], log); err != nil {
		t.Fatal(err)
	}
	execMigrationFixture(t, d, `
		INSERT INTO users(id, email, display_name, role, created_at, updated_at)
			VALUES (1, 'a@example.test', 'a', 'admin', 0, 0);
		INSERT INTO jd_areas(system_id, code_start, code_end, name, position) VALUES (1, 10, 19, 'test', 0);
		INSERT INTO jd_categories(system_id, id, area_start, code, name, system) VALUES (1, 1, 10, 11, 'inbox', 1);
		INSERT INTO documents(system_id, id, owner_id, original_blob, original_size, title, content, jd_category_id, created_at, updated_at)
			VALUES (1, 1, 1, 'sha', 1, 'Beitrag', 'Ihre Krankenversicherung', 1, 0, 0)`)
	if err := db.Migrate(t.Context(), d, stable, log); err != nil {
		t.Fatal(err)
	}
	if got := trigramHits(t, d, `"versicherung"`); got != 1 {
		t.Fatalf("existing document trigram hits = %d, want 1", got)
	}
}
