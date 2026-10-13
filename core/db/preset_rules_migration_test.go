// SPDX-License-Identifier: AGPL-3.0-or-later

package db_test

import (
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/db"
	"github.com/johnnybravo-xyz/suchi/core/db/migrations"
)

// Presets stopped seeding automations. Upgrading removes the remaining
// preset-owned rows and keeps every user-owned rule, including forks of
// edited built-ins.
func TestUpgradeRemovesPresetRulesAndKeepsUserRules(t *testing.T) {
	d, log := openSetDB(t)
	stable, err := db.LoadMigrations(migrations.FS, ".")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Migrate(t.Context(), d, stable[:3], log); err != nil {
		t.Fatal(err)
	}
	execMigrationFixture(t, d, `
		INSERT INTO automations(id, system_id, name, order_index, enabled, preset_slug, created_at, updated_at) VALUES
			(1, 1, 'Solo: file utility bills', 0, 1, 'solo', 1, 1),
			(2, 1, 'Solo: file tax documents', 1, 0, 'solo', 1, 1),
			(3, 1, 'Solo: file tax documents (edited)', 1, 1, NULL, 1, 1),
			(4, 1, 'My rule', 2, 1, NULL, 1, 1);
		INSERT INTO automation_triggers(automation_id, type, filter_content_re, created_at) VALUES
			(1, 'document_added', 'bill', 1), (3, 'document_added', 'tax', 1);
		INSERT INTO automation_actions(automation_id, order_index, kind, params_json, created_at) VALUES
			(1, 0, 'assign_title', '{"template":"Bill"}', 1), (3, 0, 'assign_title', '{"template":"Tax"}', 1)`)
	if err := db.Migrate(t.Context(), d, stable, log); err != nil {
		t.Fatal(err)
	}
	assertMigrationScalar(t, d, `SELECT count(*) FROM automations`, 2)
	assertMigrationScalar(t, d, `SELECT count(*) FROM automations WHERE id IN (3, 4)`, 2)
	assertMigrationScalar(t, d, `SELECT count(*) FROM automation_triggers WHERE automation_id NOT IN (SELECT id FROM automations)`, 0)
	assertMigrationScalar(t, d, `SELECT count(*) FROM automation_actions WHERE automation_id NOT IN (SELECT id FROM automations)`, 0)
	assertMigrationScalar(t, d, `SELECT count(*) FROM pragma_table_info('automations') WHERE name = 'preset_slug'`, 0)
}
