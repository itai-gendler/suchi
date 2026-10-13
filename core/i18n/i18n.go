// SPDX-License-Identifier: AGPL-3.0-or-later

// Package i18n loads the server-rendered UI string catalogs.
package i18n

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

//go:embed strings/*.yaml
var catalogFS embed.FS

const (
	DefaultLocale = "en"
	CookieName    = "suchi_locale"
)

// Catalog is one language's string table.
type Catalog struct {
	Locale    string
	Direction string
	Strings   map[string]string
	fallback  map[string]string
	log       *slog.Logger
}

// Registry owns every catalog embedded in this build and negotiates requests.
type Registry struct {
	catalogs map[string]*Catalog
	locales  []string
	matcher  language.Matcher
}

// NormalizeLocale converts a browser-style locale tag to a canonical lowercase key.
func NormalizeLocale(locale string) string {
	locale = strings.TrimSpace(strings.Split(locale, ",")[0])
	locale = strings.ReplaceAll(locale, "_", "-")
	locale = strings.TrimSpace(strings.Split(locale, ";")[0])
	if locale == "" || locale == "*" {
		return DefaultLocale
	}
	tag, err := language.Parse(locale)
	if err != nil {
		return DefaultLocale
	}
	return strings.ToLower(tag.String())
}

// NewRegistry validates and loads every embedded catalog once at startup.
func NewRegistry(log *slog.Logger) (*Registry, error) {
	return newRegistry(catalogFS, "strings/*.yaml", log)
}

func newRegistry(fsys fs.FS, pattern string, log *slog.Logger) (*Registry, error) {
	if log == nil {
		log = slog.Default()
	}
	paths, err := fs.Glob(fsys, pattern)
	if err != nil {
		return nil, fmt.Errorf("list catalogs: %w", err)
	}
	if len(paths) == 0 {
		return nil, errors.New("no translation catalogs available")
	}
	loaded := make(map[string]map[string]string, len(paths))
	for _, path := range paths {
		locale := strings.ToLower(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
		catalog, err := loadCatalog(fsys, path)
		if err != nil {
			return nil, err
		}
		loaded[locale] = catalog
	}
	fallback, ok := loaded[DefaultLocale]
	if !ok {
		return nil, errors.New("English fallback catalog is missing")
	}
	keys := sortedKeys(fallback)
	locales := make([]string, 0, len(loaded))
	for locale, messages := range loaded {
		if err := validateCatalog(locale, keys, fallback, messages); err != nil {
			return nil, err
		}
		locales = append(locales, locale)
	}
	sort.Strings(locales)
	if i := sort.SearchStrings(locales, DefaultLocale); i < len(locales) && locales[i] == DefaultLocale && i != 0 {
		copy(locales[1:i+1], locales[0:i])
		locales[0] = DefaultLocale
	}
	tags := make([]language.Tag, 0, len(locales))
	catalogs := make(map[string]*Catalog, len(locales))
	for _, locale := range locales {
		tag, err := language.Parse(locale)
		if err != nil {
			return nil, fmt.Errorf("parse catalog locale %q: %w", locale, err)
		}
		tags = append(tags, tag)
		direction := loaded[locale]["meta.direction"]
		if direction != "ltr" && direction != "rtl" {
			return nil, fmt.Errorf("catalog %s has invalid meta.direction %q", locale, direction)
		}
		catalogs[locale] = &Catalog{
			Locale: locale, Direction: direction, Strings: loaded[locale], fallback: fallback,
			log: log.With("component", "i18n", "locale", locale),
		}
	}
	return &Registry{catalogs: catalogs, locales: locales, matcher: language.NewMatcher(tags)}, nil
}

func loadCatalog(fsys fs.FS, path string) (map[string]string, error) {
	content, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("read catalog %s: %w", path, err)
	}
	var messages map[string]string
	if err := yaml.Unmarshal(content, &messages); err != nil {
		return nil, fmt.Errorf("parse catalog %s: %w", path, err)
	}
	if len(messages) == 0 {
		return nil, fmt.Errorf("catalog %s is empty", path)
	}
	return messages, nil
}

func validateCatalog(locale string, expectedKeys []string, english, messages map[string]string) error {
	if len(messages) != len(expectedKeys) {
		return fmt.Errorf("catalog %s has %d messages; English has %d", locale, len(messages), len(expectedKeys))
	}
	for _, key := range expectedKeys {
		value, ok := messages[key]
		if !ok || value == "" {
			return fmt.Errorf("catalog %s is missing %q", locale, key)
		}
		if fmtPlaceholders(value) != fmtPlaceholders(english[key]) {
			return fmt.Errorf("catalog %s placeholders differ for %q", locale, key)
		}
	}
	return nil
}

func sortedKeys(messages map[string]string) []string {
	keys := make([]string, 0, len(messages))
	for key := range messages {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func fmtPlaceholders(message string) string {
	var placeholders []string
	for rest := message; ; {
		start := strings.IndexByte(rest, '%')
		if start < 0 {
			break
		}
		rest = rest[start+1:]
		if strings.HasPrefix(rest, "%") {
			rest = rest[1:]
			continue
		}
		end := strings.IndexAny(rest, "vTtbcdoOqxXUeEfgGsxp")
		if end < 0 {
			placeholders = append(placeholders, "%")
			break
		}
		placeholders = append(placeholders, rest[:end+1])
		rest = rest[end+1:]
	}
	return strings.Join(placeholders, "\x00")
}

// Available reports the catalogs embedded in this build, English first.
func (r *Registry) Available() []string {
	return append([]string(nil), r.locales...)
}

// Resolve matches a locale or Accept-Language value against available catalogs.
func (r *Registry) Resolve(preference string) *Catalog {
	if r == nil {
		return nil
	}
	if preference == "" {
		return r.catalogs[DefaultLocale]
	}
	_, index := language.MatchStrings(r.matcher, preference)
	if index < 0 || index >= len(r.locales) {
		return r.catalogs[DefaultLocale]
	}
	return r.catalogs[r.locales[index]]
}

// ForRequest resolves an explicit locale cookie before Accept-Language.
func (r *Registry) ForRequest(request *http.Request) *Catalog {
	if request != nil {
		if cookie, err := request.Cookie(CookieName); err == nil && cookie.Value != "" {
			return r.Resolve(cookie.Value)
		}
		if preference := request.Header.Get("Accept-Language"); preference != "" {
			return r.Resolve(preference)
		}
	}
	return r.Resolve(DefaultLocale)
}

// Supports reports whether an exact or base locale has a catalog.
func (r *Registry) Supports(locale string) bool {
	normalized := NormalizeLocale(locale)
	if _, ok := r.catalogs[normalized]; ok {
		return true
	}
	base, _, ok := strings.Cut(normalized, "-")
	_, exists := r.catalogs[base]
	return ok && exists
}

// Lookup returns a translated string, falling back to English.
func (c *Catalog) Lookup(key string) (string, bool) {
	if value, ok := c.Strings[key]; ok {
		return value, true
	}
	value, ok := c.fallback[key]
	return value, ok
}

// Msg returns a visible marker for missing keys so catalog gaps are obvious.
func (c *Catalog) Msg(key string) string {
	if value, ok := c.Lookup(key); ok {
		return value
	}
	c.log.Warn("i18n.miss", "key", key)
	return "!" + key + "!"
}

// Msgf formats a catalog message with fmt-style arguments.
func (c *Catalog) Msgf(key string, args ...any) string {
	return fmt.Sprintf(c.Msg(key), args...)
}

// FuncMap exposes msg and formatted msg lookups to html/template.
func (c *Catalog) FuncMap() template.FuncMap {
	return template.FuncMap{
		"msg":  c.Msg,
		"msgf": c.Msgf,
	}
}
