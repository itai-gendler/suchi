// SPDX-License-Identifier: AGPL-3.0-or-later

package db_test

import (
	"io"
	"log/slog"
	"path/filepath"
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/db"
	"github.com/johnnybravo-xyz/suchi/core/db/migrations"
)

func TestShareVisibilityMigrationBackfillsAndLocksVisibility(t *testing.T) {
	migs, err := db.LoadMigrations(migrations.FS, ".")
	if err != nil {
		t.Fatal(err)
	}
	d, err := db.Open(t.Context(), filepath.Join(t.TempDir(), "share-visibility.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	if err := db.Migrate(t.Context(), d, migs[:3], log); err != nil {
		t.Fatal(err)
	}
	if _, err := d.ExecWrite(t.Context(), `
		INSERT INTO users(id,email,display_name,role,created_at,updated_at)
		VALUES(1,'owner@example.test','Owner','admin',0,0);
		INSERT INTO share_links(id,system_id,token,doc_ids_json,created_by,created_at)
		VALUES(1,1,'existing','[]',1,10);
	`); err != nil {
		t.Fatal(err)
	}
	if err := db.Migrate(t.Context(), d, migs, log); err != nil {
		t.Fatal(err)
	}

	var visibility string
	if err := d.Read.QueryRowContext(t.Context(), `SELECT visibility FROM share_links WHERE id=1`).Scan(&visibility); err != nil {
		t.Fatal(err)
	}
	if visibility != "network" {
		t.Fatalf("existing share visibility=%q, want network", visibility)
	}
	if _, err := d.ExecWrite(t.Context(), `
		INSERT INTO share_links(id,system_id,token,doc_ids_json,created_by,visibility,expires_at,created_at)
		VALUES(2,1,'public','[]',1,'public',20,10)
	`); err != nil {
		t.Fatal(err)
	}
	if _, err := d.ExecWrite(t.Context(), `UPDATE share_links SET visibility='public' WHERE id=1`); err == nil {
		t.Fatal("visibility update succeeded")
	}
	if _, err := d.ExecWrite(t.Context(), `
		INSERT OR REPLACE INTO share_links(id,system_id,token,doc_ids_json,created_by,visibility,created_at)
		VALUES(2,1,'replacement','[]',1,'network',10)
	`); err == nil {
		t.Fatal("visibility replacement succeeded")
	}
	if _, err := d.ExecWrite(t.Context(), `
		INSERT INTO share_links(system_id,token,doc_ids_json,created_by,visibility,created_at)
		VALUES(1,'invalid','[]',1,'private',10)
	`); err == nil {
		t.Fatal("invalid visibility succeeded")
	}
}
