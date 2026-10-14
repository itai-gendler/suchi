// SPDX-License-Identifier: AGPL-3.0-or-later

package exporter

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/johnnybravo-xyz/suchi/core/db"
	"github.com/johnnybravo-xyz/suchi/core/jd/presetfile"
	"github.com/johnnybravo-xyz/suchi/core/jd/systems"
)

// BuildExport reads one archive's filing tree. Automations are archive behavior
// and travel only in a complete backup.
func BuildExport(ctx context.Context, d *db.DB, systemID int64) (*presetfile.PresetFile, error) {
	tx, err := d.Read.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	pf, err := ReadTree(ctx, tx, systemID)
	if err != nil {
		return nil, err
	}
	if err := presetfile.ValidateNormalized(pf); err != nil {
		return nil, fmt.Errorf("archive cannot be exported as v1: %w", err)
	}
	return pf, nil
}

// ReadTree reads the current filing map in the caller's transaction. An empty
// database is allowed for first import; a populated tree must have System/49.
// Legacy reserved structures are diagnosed, never hidden or repaired.
func ReadTree(ctx context.Context, tx *sql.Tx, systemID int64) (*presetfile.PresetFile, error) {
	pf := &presetfile.PresetFile{
		Format: presetfile.Format, ID: "archive", Version: 1,
		Name: "Archive filing snapshot", Market: "global", Language: "und",
		Story: "Current filing structure from this archive, including local edits and merged imports. This generated snapshot is not an upstream preset or a complete archive backup.",
		Areas: []presetfile.Area{},
	}
	system, err := systems.Get(ctx, tx, systemID)
	if err != nil {
		return nil, err
	}
	pf.System, pf.Name = system.Code, system.Name
	rows, err := tx.QueryContext(ctx, `SELECT code_start, code_end, name FROM jd_areas WHERE system_id=? ORDER BY position, code_start`, systemID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a presetfile.Area
		var end int
		if err := rows.Scan(&a.Code, &end, &a.Name); err != nil {
			rows.Close()
			return nil, err
		}
		if end != a.Code+9 || a.Code < 10 || a.Code > 90 || a.Code%10 != 0 || (a.Code == 40 && a.Name != "System") {
			rows.Close()
			return nil, fmt.Errorf("legacy area %d %q is incompatible with v1; review the filing tree explicitly before import/export", a.Code, a.Name)
		}
		a.Categories = []presetfile.Category{}
		pf.Areas = append(pf.Areas, a)
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, err
	}
	rows.Close()
	var inboxID int64
	for i := range pf.Areas {
		a := &pf.Areas[i]
		rows, err := tx.QueryContext(ctx, `SELECT id, code, name, COALESCE(description,''), system FROM jd_categories WHERE system_id=? AND area_start = ? ORDER BY code`, systemID, a.Code)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var c presetfile.Category
			var id int64
			var system bool
			if err := rows.Scan(&id, &c.Code, &c.Name, &c.Description, &system); err != nil {
				rows.Close()
				return nil, err
			}
			if (a.Code == 40 && (c.Code != 49 || c.Name != "Inbox" || !system)) || (a.Code != 40 && system) {
				rows.Close()
				return nil, fmt.Errorf("legacy category %d %q conflicts with generated System/49 Inbox; review it explicitly before import/export", c.Code, c.Name)
			}
			if system {
				inboxID = id
				pf.Inbox = c.Code
			}
			a.Categories = append(a.Categories, c)
		}
		if err := rows.Err(); err != nil {
			rows.Close()
			return nil, err
		}
		rows.Close()
	}
	if len(pf.Areas) == 0 {
		return pf, nil
	}
	var configured int64
	if err := tx.QueryRowContext(ctx, `SELECT COALESCE(inbox_category_id,0) FROM jd_systems WHERE id=?`, systemID).Scan(&configured); err != nil {
		return nil, fmt.Errorf("archive Inbox pointer is unreadable: %w", err)
	}
	if configured != inboxID || inboxID == 0 {
		return nil, fmt.Errorf("archive Inbox setting must reference the generated 49 Inbox; review it before import/export")
	}
	if err := presetfile.ValidateNormalized(pf); err != nil {
		return nil, fmt.Errorf("legacy filing tree is incompatible with v1; review it explicitly: %w", err)
	}
	return pf, nil
}
