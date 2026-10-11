// SPDX-License-Identifier: AGPL-3.0-or-later

package zugferd

import (
	"io"
	"log/slog"
	"path/filepath"
	"strings"
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/db"
	"github.com/johnnybravo-xyz/suchi/core/db/migrations"
)

func TestApplyCreatesManagedInvoiceFields(t *testing.T) {
	d := openApplyTestDB(t)
	docID := seedApplyTestDocument(t, d)

	if err := Apply(t.Context(), d, docID, &Invoice{
		Number:     "INV-42",
		IssueDate:  "2026-07-31",
		Currency:   "EUR",
		TotalGross: 42.50,
	}); err != nil {
		t.Fatal(err)
	}

	var managed int
	if err := d.Read.QueryRowContext(t.Context(), `
		SELECT count(*) FROM custom_fields
		WHERE system_id=1 AND managed_by=?
	`, CustomFieldOwner).Scan(&managed); err != nil {
		t.Fatal(err)
	}
	if managed != len(invoiceFieldDefinitions) {
		t.Fatalf("managed invoice fields = %d, want %d", managed, len(invoiceFieldDefinitions))
	}
	var number string
	if err := d.Read.QueryRowContext(t.Context(), `
		SELECT v.value_text
		FROM document_custom_field_values v
		JOIN custom_fields f ON f.id=v.field_id
		WHERE v.document_id=? AND f.name='invoice_number'
	`, docID).Scan(&number); err != nil {
		t.Fatal(err)
	}
	if number != "INV-42" {
		t.Fatalf("invoice number = %q, want INV-42", number)
	}
}

func TestApplyAdoptsCompatibleFieldAndRejectsTypeConflict(t *testing.T) {
	t.Run("adopts compatible field", func(t *testing.T) {
		d := openApplyTestDB(t)
		docID := seedApplyTestDocument(t, d)
		if _, err := d.ExecWrite(t.Context(), `
			INSERT INTO custom_fields(system_id,name,data_type,created_at,updated_at)
			VALUES(1,'invoice_number','text',0,0)
		`); err != nil {
			t.Fatal(err)
		}
		if err := Apply(t.Context(), d, docID, &Invoice{Number: "INV-7"}); err != nil {
			t.Fatal(err)
		}
		var managedBy string
		if err := d.Read.QueryRowContext(t.Context(), `
			SELECT managed_by FROM custom_fields
			WHERE system_id=1 AND name='invoice_number'
		`).Scan(&managedBy); err != nil {
			t.Fatal(err)
		}
		if managedBy != CustomFieldOwner {
			t.Fatalf("managed_by = %q, want %q", managedBy, CustomFieldOwner)
		}
	})

	t.Run("rejects conflicting type", func(t *testing.T) {
		d := openApplyTestDB(t)
		docID := seedApplyTestDocument(t, d)
		if _, err := d.ExecWrite(t.Context(), `
			INSERT INTO custom_fields(system_id,name,data_type,created_at,updated_at)
			VALUES(1,'invoice_date','text',0,0)
		`); err != nil {
			t.Fatal(err)
		}
		err := Apply(t.Context(), d, docID, &Invoice{IssueDate: "2026-07-31"})
		if err == nil || !strings.Contains(err.Error(), `requires "date"`) {
			t.Fatalf("Apply conflict error = %v", err)
		}
		var created int
		if err := d.Read.QueryRowContext(t.Context(), `
			SELECT count(*) FROM custom_fields WHERE managed_by=?
		`, CustomFieldOwner).Scan(&created); err != nil {
			t.Fatal(err)
		}
		if created != 0 {
			t.Fatalf("failed Apply committed %d managed fields", created)
		}
	})
}

func openApplyTestDB(t *testing.T) *db.DB {
	t.Helper()
	d, err := db.Open(t.Context(), filepath.Join(t.TempDir(), "zugferd.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = d.Close() })
	migs, err := db.LoadMigrations(migrations.FS, ".")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Migrate(t.Context(), d, migs,
		slog.New(slog.NewTextHandler(io.Discard, nil))); err != nil {
		t.Fatal(err)
	}
	return d
}

func seedApplyTestDocument(t *testing.T, d *db.DB) int64 {
	t.Helper()
	if _, err := d.ExecWrite(t.Context(), `
		INSERT INTO users(id,email,display_name,role,created_at,updated_at)
		VALUES(1,'owner@example.test','Owner','admin',0,0);
		INSERT INTO jd_areas(system_id,code_start,code_end,name,position)
		VALUES(1,0,9,'System',0);
		INSERT INTO jd_categories(id,system_id,area_start,code,name,system)
		VALUES(1,1,0,1,'Inbox',1)
	`); err != nil {
		t.Fatal(err)
	}
	result, err := d.ExecWrite(t.Context(), `
		INSERT INTO documents(
			system_id,owner_id,original_blob,original_size,title,mime_type,
			jd_category_id,created_at,added_at,updated_at
		)
		VALUES(1,1,'original',1,'Invoice','application/pdf',1,0,0,0)
	`)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	return id
}
