package generator

import (
	"fmt"
	"math/rand"
)

func GenerateCNPJ(masked bool) string {
	cnpj := make([]int, 14)

	for i := 0; i < 12; i++ {
		cnpj[i] = rand.Intn(10)
	}

	cnpj[12] = calculateCNPJDigit(cnpj[:12], []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	cnpj[13] = calculateCNPJDigit(cnpj[:13], []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})

	if masked {
		return fmt.Sprintf("%d%d.%d%d%d.%d%d%d/%d%d%d%d-%d%d",
			cnpj[0], cnpj[1],
			cnpj[2], cnpj[3], cnpj[4],
			cnpj[5], cnpj[6], cnpj[7],
			cnpj[8], cnpj[9], cnpj[10], cnpj[11],
			cnpj[12], cnpj[13])
	}

	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d",
		cnpj[0], cnpj[1], cnpj[2], cnpj[3], cnpj[4], cnpj[5], cnpj[6],
		cnpj[7], cnpj[8], cnpj[9], cnpj[10], cnpj[11], cnpj[12], cnpj[13])
}

func calculateCNPJDigit(digits []int, weights []int) int {
	sum := 0

	for i, digit := range digits {
		sum += digit * weights[i]
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
