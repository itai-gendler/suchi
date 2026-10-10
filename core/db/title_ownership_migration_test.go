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

func TestTitleOwnershipMigrationBackfillsAndTracksOwnerChanges(t *testing.T) {
	migs, err := db.LoadMigrations(migrations.FS, ".")
	if err != nil {
		t.Fatal(err)
	}
	d, err := db.Open(t.Context(), filepath.Join(t.TempDir(), "title-ownership.db"))
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
		INSERT INTO jd_areas(system_id,code_start,code_end,name,position)
		VALUES(1,0,9,'System',0);
		INSERT INTO jd_categories(id,system_id,area_start,code,name,system)
		VALUES(1,1,0,1,'Inbox',1);
		INSERT INTO documents(id,system_id,owner_id,original_blob,original_size,title,jd_category_id,created_at,updated_at) VALUES
			(10,1,1,'filename',1,'scan',1,0,0),
			(11,1,1,'untitled',1,'Untitled',1,0,0),
			(12,1,1,'custom',1,'Quarterly scan',1,0,0),
			(13,1,1,'paperless',1,'paperless',1,0,0),
			(14,1,1,'watched',1,'report',1,0,0);
		INSERT INTO document_sources(id,document_id,kind,label,detail,observed_at) VALUES
			(1,10,'upload','Browser','incoming/scan.PDF',1),
			(2,12,'upload','Browser','scan.pdf',1),
			(3,13,'import','Paperless-ngx import','paperless.pdf',1),
			(4,14,'watched_folder','Inbox','archive/report.docx',1);
	`); err != nil {
		t.Fatal(err)
	}
	if err := db.Migrate(t.Context(), d, migs, log); err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		id   int64
		want string
	}{
		{10, "auto"},
		{11, "auto"},
		{12, "human"},
		{13, "human"},
		{14, "auto"},
	} {
		var got string
		if err := d.Read.QueryRowContext(t.Context(), `SELECT title_owner FROM documents WHERE id=?`, tc.id).Scan(&got); err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Errorf("document %d title_owner=%q, want %q", tc.id, got, tc.want)
		}
	}

	if _, err := d.ExecWrite(t.Context(), `UPDATE documents SET title_owner='human' WHERE id=10`); err != nil {
		t.Fatal(err)
	}
	var revision int
	if err := d.Read.QueryRowContext(t.Context(), `SELECT title_revision FROM documents WHERE id=10`).Scan(&revision); err != nil {
		t.Fatal(err)
	}
	if revision != 1 {
		t.Fatalf("owner-only update title_revision=%d, want 1", revision)
	}
	if _, err := d.ExecWrite(t.Context(), `UPDATE documents SET title='renamed',title_owner='auto' WHERE id=10`); err != nil {
		t.Fatal(err)
	}
	if err := d.Read.QueryRowContext(t.Context(), `SELECT title_revision FROM documents WHERE id=10`).Scan(&revision); err != nil {
		t.Fatal(err)
	}
	if revision != 2 {
		t.Fatalf("title and owner update title_revision=%d, want 2", revision)
	}
	if _, err := d.ExecWrite(t.Context(), `UPDATE documents SET title_owner='model' WHERE id=10`); err == nil {
		t.Fatal("invalid title owner accepted")
	}
}
