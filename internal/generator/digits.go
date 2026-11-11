// Package generator creates and validates Brazilian CPF and CNPJ numbers.
package generator

import (
	"math/rand/v2"
	"strings"
)

const (
	digitChars        = "0123456789"
	alphanumericChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// randomChars returns n random characters from charset that are not all equal.
// Numbers made of a single repeated character (e.g. 111.111.111-11) pass the
// check-digit math but are rejected by real validators, so they are skipped.
func randomChars(n int, charset string) []byte {
	for {
		chars := make([]byte, n)
		for i := range chars {
			chars[i] = charset[rand.IntN(len(charset))]
		}
		if !allEqual(chars) {
			return chars
		}
	}
}

// charValue is the value a character has in the check-digit calculation:
// its ASCII code minus 48. Digits map to 0-9 and letters A-Z to 17-42, as
// defined by the Receita Federal for the alphanumeric CNPJ.
func charValue(c byte) int {
	return int(c) - '0'
}

// checkDigit computes a mod-11 check digit over chars using the given weights.
func checkDigit(chars []byte, weights []int) byte {
	sum := 0
	for i, c := range chars {
		sum += charValue(c) * weights[i]
	}
	if r := sum % 11; r >= 2 {
		return byte('0' + 11 - r)
	}
	return '0'
}

func allEqual(chars []byte) bool {
	for _, c := range chars[1:] {
		if c != chars[0] {
			return false
		}
	}
	return true
}

// format writes chars into a mask, replacing each '#' with the next character.
func format(chars []byte, mask string) string {
	var b strings.Builder
	b.Grow(len(mask))
	i := 0
	for _, m := range []byte(mask) {
		if m == '#' {
			b.WriteByte(chars[i])
			i++
		} else {
			b.WriteByte(m)
		}
	}
	return b.String()
}

// parse removes the mask punctuation ('.', '-', '/') from s. It returns
// ok=true only if exactly n characters are left, all from charset, and the
// last two (the check digits) are digits.
func parse(s string, n int, charset string) (chars []byte, ok bool) {
	chars = make([]byte, 0, n)
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case strings.IndexByte(charset, c) >= 0:
			chars = append(chars, c)
		case c == '.' || c == '-' || c == '/':
		default:
			return nil, false
		}
	}
	if len(chars) != n {
		return nil, false
	}
	for _, c := range chars[n-2:] {
		if c < '0' || c > '9' {
			return nil, false
		}
	}
	return chars, true
}
