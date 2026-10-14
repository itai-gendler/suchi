// SPDX-License-Identifier: AGPL-3.0-or-later

package bundle_test

import (
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/importer/bundle"
)

// Money arrives as text with a currency code; it is stored as a number so
// money filters and display work. Unreadable text is kept as text.
func TestMonetaryValuesImportAsNumbers(t *testing.T) {
	for in, want := range map[string]string{
		"EUR1234.56": "1234.56",
		"1234.56":    "1234.56",
		"CHF0.50":    "0.5",
		"not money":  `"not money"`,
	} {
		value := in
		got := string(bundle.CustomFieldInstance{ValueMonetary: &value}.EffectiveValue())
		if got != want {
			t.Errorf("EffectiveValue(%q) = %s, want %s", in, got, want)
		}
	}
}
