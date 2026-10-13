// SPDX-License-Identifier: AGPL-3.0-or-later

package importer_test

import (
	"io"
	"log/slog"
	"path/filepath"
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/db"
	migrations "github.com/johnnybravo-xyz/suchi/core/db/migrations"
	"github.com/johnnybravo-xyz/suchi/core/jd/importer"
	"github.com/johnnybravo-xyz/suchi/core/jd/presetfile"
)

const smallTaxonomy = `format = "suchi-taxonomy/v1"
id = "smalltest"
version = 1
name = "Small test"
market = "global"
language = "en"
story = "File utility invoices while keeping unrelated paperwork in Inbox."
[[areas]]
code = 10
name = "Life"
[[areas.categories]]
code = 11
name = "Bills"
description = "Utility bills."
`

func smallPreset(t *testing.T) *presetfile.PresetFile {
	t.Helper()
	pf, err := presetfile.Parse([]byte(smallTaxonomy), presetfile.FormatTOML)
	if err != nil {
		t.Fatal(err)
	}
	return pf
}
func logger() *slog.Logger { return slog.New(slog.NewTextHandler(io.Discard, nil)) }
func openTestDB(t *testing.T) *db.DB {
	t.Helper()
	d, err := db.Open(t.Context(), filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = d.Close() })
	migs, err := db.LoadMigrations(migrations.FS, ".")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Migrate(t.Context(), d, migs, logger()); err != nil {
		t.Fatal(err)
	}
	return d
}
func apply(t *testing.T, d *db.DB, pf *presetfile.PresetFile, opts importer.Options) *importer.Diff {
	t.Helper()
	diff, err := importer.ImportForDB(t.Context(), d, logger(), pf, opts)
	if err != nil {
		t.Fatal(err)
	}
	return diff
}
func count(t *testing.T, d *db.DB, query string, args ...any) int {
	t.Helper()
	var n int
	if err := d.Read.QueryRowContext(t.Context(), query, args...).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}
func document(t *testing.T, d *db.DB, title, content string) int64 {
	t.Helper()
	if _, err := d.Write.ExecContext(t.Context(), `INSERT INTO users(id,email,display_name,role,created_at,updated_at) VALUES(1,'test@example.test','Test','admin',0,0) ON CONFLICT DO NOTHING`); err != nil {
		t.Fatal(err)
	}
	res, err := d.Write.ExecContext(t.Context(), `INSERT INTO documents(system_id,owner_id,title,content,original_blob,original_size,jd_category_id,created_at,added_at,updated_at) VALUES(1,1,?,?,?,1,(SELECT id FROM jd_categories WHERE system_id=1 AND system=1),0,0,0)`, title, content, title)
	if err != nil {
		t.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	return id
}
