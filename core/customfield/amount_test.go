// SPDX-License-Identifier: AGPL-3.0-or-later

package customfield

import "testing"

func TestParseAmount(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want float64
	}{
		{"EUR1234.56", 1234.56},
		{"1234.56", 1234.56},
		{"1.234,56 €", 1234.56},
		{"€1,234.56", 1234.56},
		{"$ 1,234,567.89", 1234567.89},
		{"1'234.50 CHF", 1234.50},
		{"1 234,56", 1234.56},
		{"1 234,56 €", 1234.56},
		{"12,5", 12.5},
		{"1.234", 1234},
		{"1.234.567", 1234567},
		{"0,99", 0.99},
		{"(12.00)", -12},
		{"-5", -5},
		{"€-5,00", -5},
		{"₹ 1,00,000.00", 100000},
	} {
		got, err := ParseAmount(tc.in)
		if err != nil || got != tc.want {
			t.Errorf("ParseAmount(%q) = %v, %v; want %v", tc.in, got, err, tc.want)
		}
	}
	for _, in := range []string{"", "abc", "EUR", "1,2,3.4.5", "12-34", "12 EUR 34", "1/2"} {
		if got, err := ParseAmount(in); err == nil {
			t.Errorf("ParseAmount(%q) = %v, want error", in, got)
		}
	}
}
