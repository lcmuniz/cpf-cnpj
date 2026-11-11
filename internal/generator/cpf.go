package generator

const cpfMask = "###.###.###-##"

// GenerateCPF returns a random valid CPF. If masked is true it is formatted
// as XXX.XXX.XXX-XX, otherwise it is returned as 11 plain digits.
func GenerateCPF(masked bool) string {
	cpf := append(randomChars(9, digitChars), 0, 0)
	cpf[9] = checkDigit(cpf[:9], cpfWeights(9))
	cpf[10] = checkDigit(cpf[:10], cpfWeights(10))

	if masked {
		return format(cpf, cpfMask)
	}
	return string(cpf)
}

// IsValidCPF reports whether s is a valid CPF, with or without the mask.
func IsValidCPF(s string) bool {
	d, ok := parse(s, 11, digitChars)
	if !ok || allEqual(d) {
		return false
	}
	return d[9] == checkDigit(d[:9], cpfWeights(9)) &&
		d[10] == checkDigit(d[:10], cpfWeights(10))
}

// cpfWeights returns the descending weights n+1 ... 2 used by the CPF algorithm.
func cpfWeights(n int) []int {
	w := make([]int, n)
	for i := range w {
		w[i] = n + 1 - i
	}
	return w
}
