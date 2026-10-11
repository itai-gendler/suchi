// SPDX-License-Identifier: AGPL-3.0-or-later

// Apply glue: take a parsed Invoice + a doc id and write the
// custom-field values. Kept in this package (rather than post-ingest)
// so post-ingest stays thin and the ZUGFeRD shape stays owned in one
// place.

package zugferd

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/johnnybravo-xyz/suchi/core/db"
)

// CustomFieldOwner identifies definitions maintained by e-invoice extraction.
const CustomFieldOwner = "zugferd"

// invoiceFieldDefinitions are the custom-field keys owned by e-invoice
// extraction.
var invoiceFieldDefinitions = [...]struct {
	Name     string
	DataType string
}{
	{"invoice_number", "text"},
	{"invoice_date", "date"},
	{"invoice_currency", "text"},
	{"invoice_total_gross", "monetary"},
	{"invoice_total_net", "monetary"},
	{"invoice_seller", "text"},
	{"invoice_seller_vat", "text"},
	{"invoice_buyer", "text"},
}

// IsInvoiceFieldName reports whether name belongs to the e-invoice extraction
// contract. Names use the same case-sensitive identity as custom_fields.
func IsInvoiceFieldName(name string) bool {
	name = strings.TrimSpace(name)
	for _, field := range invoiceFieldDefinitions {
		if name == field.Name {
			return true
		}
	}
	return false
}

// Apply writes the invoice into document_custom_field_values,
// creating the custom_fields rows if this is the first invoice on the
// instance. Idempotent: repeated calls on the same doc overwrite the
// existing values (ON CONFLICT DO UPDATE via a DELETE-then-INSERT
// pattern within the tx).
func Apply(ctx context.Context, d *db.DB, docID int64, inv *Invoice) error {
	if inv == nil {
		return nil
	}
	return d.WriteTx(ctx, func(tx *sql.Tx) error {
		var systemID int64
		if err := tx.QueryRowContext(ctx, `SELECT system_id FROM documents WHERE id = ?`, docID).Scan(&systemID); err != nil {
			return err
		}
		fieldIDs, err := ensureFields(ctx, tx, systemID)
		if err != nil {
			return err
		}
		vals := map[string]any{
			"invoice_number":      inv.Number,
			"invoice_date":        inv.IssueDate, // stored via issue-date parser below
			"invoice_currency":    inv.Currency,
			"invoice_total_gross": inv.TotalGross,
			"invoice_total_net":   inv.TotalNet,
			"invoice_seller":      inv.SellerName,
			"invoice_seller_vat":  inv.SellerVAT,
			"invoice_buyer":       inv.BuyerName,
		}
		for name, v := range vals {
			fid := fieldIDs[name]
			if fid == 0 {
				continue
			}
			if _, err := tx.ExecContext(ctx,
				`DELETE FROM document_custom_field_values
				 WHERE document_id = ? AND field_id = ?`,
				docID, fid); err != nil {
				return err
			}
			if err := writeVal(ctx, tx, docID, fid, name, v); err != nil {
				return err
			}
		}
		return nil
	})
}

// ensureFields creates missing invoice_* rows, adopts compatible definitions
// created before ownership was tracked, and returns the name→id lookup.
func ensureFields(ctx context.Context, tx *sql.Tx, systemID int64) (map[string]int64, error) {
	out := make(map[string]int64, len(invoiceFieldDefinitions))
	now := time.Now().Unix()
	for _, field := range invoiceFieldDefinitions {
		var (
			id        int64
			dataType  string
			managedBy string
		)
		err := tx.QueryRowContext(ctx, `
			SELECT id, data_type, managed_by
			FROM custom_fields
			WHERE system_id = ? AND name = ?
		`, systemID, field.Name).Scan(&id, &dataType, &managedBy)
		if err == sql.ErrNoRows {
			res, err := tx.ExecContext(ctx, `
				INSERT INTO custom_fields(
					system_id, name, data_type, managed_by, created_at, updated_at
				)
				VALUES (?, ?, ?, ?, ?, ?)
			`, systemID, field.Name, field.DataType, CustomFieldOwner, now, now)
			if err != nil {
				return nil, err
			}
			id, err = res.LastInsertId()
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		} else {
			if dataType != field.DataType {
				return nil, fmt.Errorf(
					"reserved custom field %q has type %q; e-invoice extraction requires %q",
					field.Name, dataType, field.DataType,
				)
			}
			if managedBy != "" && managedBy != CustomFieldOwner {
				return nil, fmt.Errorf("custom field %q is managed by %q", field.Name, managedBy)
			}
			if managedBy == "" {
				if _, err := tx.ExecContext(ctx, `
					UPDATE custom_fields
					SET managed_by = ?, updated_at = ?
					WHERE id = ? AND managed_by = ''
				`, CustomFieldOwner, now, id); err != nil {
					return nil, err
				}
			}
		}
		out[field.Name] = id
	}
	return out, nil
}

// writeVal routes v into the right typed column. Skips empties so we
// don't clutter the values table with placeholder rows.
func writeVal(ctx context.Context, tx *sql.Tx, docID, fieldID int64, name string, v any) error {
	switch x := v.(type) {
	case string:
		if x == "" {
			return nil
		}
		if name == "invoice_date" {
			ts := parseISODate(x)
			if ts == 0 {
				return nil
			}
			_, err := tx.ExecContext(ctx, `
				INSERT INTO document_custom_field_values(document_id, field_id, value_date)
				VALUES (?, ?, ?)
			`, docID, fieldID, ts)
			return err
		}
		_, err := tx.ExecContext(ctx, `
			INSERT INTO document_custom_field_values(document_id, field_id, value_text)
			VALUES (?, ?, ?)
		`, docID, fieldID, x)
		return err
	case float64:
		if x == 0 {
			return nil
		}
		_, err := tx.ExecContext(ctx, `
			INSERT INTO document_custom_field_values(document_id, field_id, value_number)
			VALUES (?, ?, ?)
		`, docID, fieldID, x)
		return err
	}
	return nil
}

// parseISODate turns "YYYY-MM-DD" into a unix epoch (UTC midnight).
// Returns 0 on any parse failure so writeVal can skip.
func parseISODate(s string) int64 {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return 0
	}
	return t.UTC().Unix()
}
