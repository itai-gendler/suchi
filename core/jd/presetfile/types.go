// SPDX-License-Identifier: AGPL-3.0-or-later

// Package presetfile validates raw HuML and TOML taxonomy files before
// normalizing them into the application's filing tree.
package presetfile

const Format = "suchi-taxonomy/v1"

// PresetFile is the runtime tree returned by Parse. Inbox is generated, never authored.
type PresetFile struct {
	Format     string     `huml:"format" toml:"format" json:"format"`
	System     string     `huml:"system,omitempty" toml:"system,omitempty" json:"system,omitempty"`
	ID         string     `huml:"id" toml:"id" json:"id"`
	Version    int        `huml:"version" toml:"version" json:"version"`
	Name       string     `huml:"name" toml:"name" json:"name"`
	Market     string     `huml:"market" toml:"market" json:"market"`
	Language   string     `huml:"language" toml:"language" json:"language"`
	Maintainer string     `huml:"maintainer,omitempty" toml:"maintainer,omitempty" json:"maintainer,omitempty"`
	License    string     `huml:"license,omitempty" toml:"license,omitempty" json:"license,omitempty"`
	Flat       bool       `huml:"flat,omitempty" toml:"flat,omitempty" json:"flat,omitempty"`
	Inbox      int        `huml:"-" toml:"-" json:"-"`
	Story      string     `huml:"story" toml:"story" json:"story"`
	Areas      []Area     `huml:"areas" toml:"areas" json:"areas"`
	Categories []Category `huml:"categories,omitempty" toml:"categories,omitempty" json:"categories,omitempty"`
}

type Area struct {
	Code       int        `huml:"code" toml:"code" json:"code"`
	Name       string     `huml:"name" toml:"name" json:"name"`
	Categories []Category `huml:"categories" toml:"categories" json:"categories"`
}

type Category struct {
	Code        int    `huml:"code" toml:"code" json:"code"`
	Name        string `huml:"name" toml:"name" json:"name"`
	Description string `huml:"description,omitempty" toml:"description,omitempty" json:"description,omitempty"`
}
