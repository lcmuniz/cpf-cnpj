package generator

import "strings"

const cnpjMask = "##.###.###/####-##"

var (
	cnpjWeights1 = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjWeights2 = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// GenerateCNPJ returns a random valid numeric CNPJ. If masked is true it is
// formatted as XX.XXX.XXX/XXXX-XX, otherwise it is returned as 14 plain digits.
func GenerateCNPJ(masked bool) string {
	return generateCNPJ(randomChars(12, digitChars), masked)
}

// GenerateAlphanumericCNPJ returns a random valid CNPJ in the alphanumeric
// format issued from July 2026: the first 12 characters may be digits or
// uppercase letters (at least one is a letter) and the last 2 are numeric
// check digits.
//
// The Receita Federal also excludes some letter combinations (to avoid
// offensive words) but does not publish the list, so a generated value may
// in rare cases be one that will never be issued. It is still valid for
// check-digit purposes.
func GenerateAlphanumericCNPJ(masked bool) string {
	for {
		base := randomChars(12, alphanumericChars)
		if strings.ContainsAny(string(base), alphanumericChars[10:]) {
			return generateCNPJ(base, masked)
		}
	}
}

func generateCNPJ(base []byte, masked bool) string {
	cnpj := append(base, 0, 0)
	cnpj[12] = checkDigit(cnpj[:12], cnpjWeights1)
	cnpj[13] = checkDigit(cnpj[:13], cnpjWeights2)

	if masked {
		return format(cnpj, cnpjMask)
	}
	return string(cnpj)
}

// IsValidCNPJ reports whether s is a valid CNPJ, with or without the mask.
// It accepts both the numeric format and the alphanumeric one (uppercase
// letters in the first 12 positions, numeric check digits).
func IsValidCNPJ(s string) bool {
	c, ok := parse(s, 14, alphanumericChars)
	if !ok || allEqual(c) {
		return false
	}
	return c[12] == checkDigit(c[:12], cnpjWeights1) &&
		c[13] == checkDigit(c[:13], cnpjWeights2)
}
