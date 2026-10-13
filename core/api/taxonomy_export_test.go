// SPDX-License-Identifier: AGPL-3.0-or-later

package api

import (
	"io"
	"log/slog"
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/jd/importer"
	"github.com/johnnybravo-xyz/suchi/core/jd/presetfile"
	"github.com/johnnybravo-xyz/suchi/core/taxonomy/exporter"
)

const portableTaxonomy = `format = "suchi-taxonomy/v1"
id = "utility-files"
version = 7
name = "Utility files"
market = "global"
language = "en"
maintainer = "Source author"
license = "CC0-1.0"
story = "File utility invoices."
[[areas]]
code = 10
name = "Life admin"
[[areas.categories]]
code = 11
name = "Bills"
description = "Invoices for household utilities."
`

func TestTaxonomyExportRoundTripsTree(t *testing.T) {
	ctx := t.Context()
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	source := openTestDB(t)
	pf, err := presetfile.Parse([]byte(portableTaxonomy), presetfile.FormatTOML)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := importer.ImportForDB(ctx, source, log, pf, importer.Options{}); err != nil {
		t.Fatal(err)
	}
	exported, err := exporter.BuildExport(ctx, source, 1)
	if err != nil {
		t.Fatal(err)
	}
	if exported.ID != "archive" || exported.Version != 1 || exported.Maintainer != "" || exported.License != "" {
		t.Fatalf("snapshot falsely claims source identity: %+v", exported)
	}
	for _, format := range []presetfile.SerFormat{presetfile.FormatHuML, presetfile.FormatTOML} {
		t.Run(string(format), func(t *testing.T) {
			data, err := presetfile.Marshal(exported, format)
			if err != nil {
				t.Fatal(err)
			}
			roundTrip, err := presetfile.Parse(data, format)
			if err != nil {
				t.Fatal(err)
			}
			dest := openTestDB(t)
			if _, err := importer.ImportForDB(ctx, dest, log, roundTrip, importer.Options{}); err != nil {
				t.Fatal(err)
			}
			var description string
			if err := dest.Read.QueryRow(`SELECT description FROM jd_categories WHERE code=11 AND name='Bills'`).Scan(&description); err != nil {
				t.Fatal(err)
			}
			if description != "Invoices for household utilities." {
				t.Fatalf("round trip changed category: %q", description)
			}
		})
	}
}

func TestTaxonomyAuthoringFormat(t *testing.T) {
	for _, value := range []string{"huml", "HuML", "toml", " TOML "} {
		if _, ok := taxonomyAuthoringFormat(value, ""); !ok {
			t.Errorf("taxonomyAuthoringFormat(%q) rejected", value)
		}
	}
	if got, ok := taxonomyAuthoringFormat("", "huml"); !ok || got != "huml" {
		t.Fatalf("fallback = %q, %v", got, ok)
	}
	for _, value := range []string{"yaml", "json"} {
		if _, ok := taxonomyAuthoringFormat(value, ""); ok {
			t.Errorf("taxonomyAuthoringFormat(%q) accepted", value)
		}
	}
}
