// SPDX-License-Identifier: AGPL-3.0-or-later

package i18n

import (
	"io"
	"io/fs"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"
	"testing/fstest"
)

func TestNormalizeLocale(t *testing.T) {
	tests := map[string]string{
		"":               "en",
		" EN_us ":        "en-us",
		"pt-BR":          "pt-br",
		"de-DE,de;q=0.9": "de-de",
	}
	for input, want := range tests {
		if got := NormalizeLocale(input); got != want {
			t.Errorf("NormalizeLocale(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestRegistryResolvesAvailableFallback(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	registry, err := NewRegistry(log)
	if err != nil {
		t.Fatal(err)
	}
	if got := registry.Available(); len(got) != 1 || got[0] != "en" {
		t.Fatalf("Available() = %v, want [en]", got)
	}
	for _, locale := range []string{"en-US", "fr-FR"} {
		catalog := registry.Resolve(locale)
		if catalog.Locale != "en" {
			t.Errorf("Resolve(%q) locale = %q, want en", locale, catalog.Locale)
		}
		if catalog.Direction != "ltr" {
			t.Errorf("Resolve(%q) direction = %q, want ltr", locale, catalog.Direction)
		}
		if got := catalog.Msg("auth.login.title"); got != "Sign in" {
			t.Errorf("Resolve(%q) auth.login.title = %q", locale, got)
		}
	}
}

func TestRegistryNegotiatesRequest(t *testing.T) {
	registry, err := NewRegistry(nil)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Accept-Language", "fr-FR, en;q=0.8")
	if got := registry.ForRequest(req).Locale; got != "en" {
		t.Fatalf("ForRequest() = %q, want en", got)
	}
}

func TestRegistryDiscoversCompleteGermanFixture(t *testing.T) {
	english, err := fs.ReadFile(catalogFS, "strings/en.yaml")
	if err != nil {
		t.Fatal(err)
	}
	german, err := fs.ReadFile(os.DirFS("."), "testdata/strings/de.yaml")
	if err != nil {
		t.Fatal(err)
	}
	registry, err := newRegistry(fstest.MapFS{
		"catalogs/en.yaml": {Data: english},
		"catalogs/de.yaml": {Data: german},
	}, "catalogs/*.yaml", nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := registry.Available(); len(got) != 2 || got[0] != "en" || got[1] != "de" {
		t.Fatalf("Available() = %v, want [en de]", got)
	}
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Accept-Language", "de-DE,de;q=0.9,en;q=0.8")
	catalog := registry.ForRequest(req)
	if catalog.Locale != "de" || catalog.Direction != "ltr" || catalog.Msg("auth.login.title") != "Anmelden" {
		t.Fatalf("German request catalog = locale=%q dir=%q title=%q", catalog.Locale, catalog.Direction, catalog.Msg("auth.login.title"))
	}
}

func TestLookupAndMissingMarker(t *testing.T) {
	registry, err := NewRegistry(nil)
	if err != nil {
		t.Fatal(err)
	}
	catalog := registry.Resolve("en")
	if got, ok := catalog.Lookup("auth.login.title"); !ok || got != "Sign in" {
		t.Fatalf("Lookup(auth.login.title) = %q, %v", got, ok)
	}
	if got := catalog.Msg("missing.key"); got != "!missing.key!" {
		t.Fatalf("Msg(missing.key) = %q", got)
	}
}
