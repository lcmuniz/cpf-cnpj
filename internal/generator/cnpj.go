package generator

const cnpjMask = "##.###.###/####-##"

var (
	cnpjWeights1 = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjWeights2 = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// GenerateCNPJ returns a random valid numeric CNPJ. If masked is true it is
// formatted as XX.XXX.XXX/XXXX-XX, otherwise it is returned as 14 plain digits.
func GenerateCNPJ(masked bool) string {
	cnpj := append(randomDigits(12), 0, 0)
	cnpj[12] = checkDigit(cnpj[:12], cnpjWeights1)
	cnpj[13] = checkDigit(cnpj[:13], cnpjWeights2)

	if masked {
		return format(cnpj, cnpjMask)
	}
	return format(cnpj, "##############")
}

// IsValidCNPJ reports whether s is a valid numeric CNPJ, with or without the mask.
func IsValidCNPJ(s string) bool {
	d, ok := parseDigits(s, 14)
	if !ok || allEqual(d) {
		return false
	}
	return d[12] == checkDigit(d[:12], cnpjWeights1) &&
		d[13] == checkDigit(d[:13], cnpjWeights2)
}
