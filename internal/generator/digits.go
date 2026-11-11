// Package generator creates and validates Brazilian CPF and CNPJ numbers.
package generator

import (
	"math/rand/v2"
	"strings"
)

// randomDigits returns n random decimal digits that are not all equal.
// Numbers made of a single repeated digit (e.g. 111.111.111-11) pass the
// check-digit math but are rejected by real validators, so they are skipped.
func randomDigits(n int) []int {
	for {
		digits := make([]int, n)
		for i := range digits {
			digits[i] = rand.IntN(10)
		}
		if !allEqual(digits) {
			return digits
		}
	}
}

// checkDigit computes a mod-11 check digit using the given weights.
func checkDigit(digits, weights []int) int {
	sum := 0
	for i, d := range digits {
		sum += d * weights[i]
	}
	if r := sum % 11; r >= 2 {
		return 11 - r
	}
	return 0
}

func allEqual(digits []int) bool {
	for _, d := range digits[1:] {
		if d != digits[0] {
			return false
		}
	}
	return true
}

// format writes digits into a mask, replacing each '#' with the next digit.
func format(digits []int, mask string) string {
	var b strings.Builder
	b.Grow(len(mask))
	i := 0
	for _, c := range mask {
		if c == '#' {
			b.WriteByte(byte('0' + digits[i]))
			i++
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

// parseDigits strips the punctuation allowed in a mask and returns the digits.
// It returns ok=false if the input contains any other character or does not
// have exactly n digits.
func parseDigits(s string, n int) (digits []int, ok bool) {
	digits = make([]int, 0, n)
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
			digits = append(digits, int(c-'0'))
		case c == '.' || c == '-' || c == '/':
		default:
			return nil, false
		}
	}
	return digits, len(digits) == n
}
