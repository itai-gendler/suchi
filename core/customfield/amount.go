// SPDX-License-Identifier: AGPL-3.0-or-later

package customfield

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// ParseAmount reads a money value written with an optional currency code or
// symbol and local digit grouping: "EUR1234.56", "1.234,56 €", "€1,234.56",
// "1'234.50", "1 234,56", "(12.00)". When both '.' and ',' appear, the last
// one is the decimal separator. A single separator followed by exactly three
// digits groups thousands ("1.234" is 1234); otherwise it is the decimal
// separator ("12,5"). Anything else is an error rather than a guess.
func ParseAmount(text string) (float64, error) {
	s := strings.TrimSpace(text)
	negative := false
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		negative, s = true, s[1:len(s)-1]
	}
	var number strings.Builder
	started, ended := false, false
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9', (r == '.' || r == ',') && started:
			if ended {
				return 0, fmt.Errorf("amount %q has text inside the number", text)
			}
			started = true
			number.WriteRune(r)
		case r == '-' && !started && !negative:
			negative = true
		case unicode.IsSpace(r) || r == '\'' || r == '’':
			// Digit grouping inside the number, padding around it.
		case unicode.IsLetter(r) || unicode.Is(unicode.Sc, r):
			// Currency codes and symbols sit before or after the number.
			ended = started
		default:
			return 0, fmt.Errorf("amount %q has unexpected %q", text, r)
		}
	}
	digits := number.String()
	if digits == "" {
		return 0, fmt.Errorf("amount %q has no digits", text)
	}
	dot, comma := strings.LastIndexByte(digits, '.'), strings.LastIndexByte(digits, ',')
	switch {
	case dot >= 0 && comma >= 0:
		decimal, group := ".", ","
		if comma > dot {
			decimal, group = ",", "."
		}
		digits = strings.ReplaceAll(strings.ReplaceAll(digits, group, ""), decimal, ".")
	case dot >= 0 || comma >= 0:
		separator := "."
		if comma >= 0 {
			separator = ","
		}
		last := strings.LastIndex(digits, separator)
		if strings.Count(digits, separator) > 1 || len(digits)-last-1 == 3 {
			digits = strings.ReplaceAll(digits, separator, "")
		} else {
			digits = strings.Replace(digits, separator, ".", 1)
		}
	}
	value, err := strconv.ParseFloat(digits, 64)
	if err != nil {
		return 0, fmt.Errorf("amount %q is not a number", text)
	}
	if negative {
		value = -value
	}
	return value, nil
}
