package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateCPF(masked bool) string {
	cpf := make([]int, 11)

	for i := 0; i < 9; i++ {
		cpf[i] = rand.Intn(10)
	}

	cpf[9] = calculateCPFDigit(cpf[:9])
	cpf[10] = calculateCPFDigit(cpf[:10])

	if masked {
		return fmt.Sprintf("%d%d%d.%d%d%d.%d%d%d-%d%d",
			cpf[0], cpf[1], cpf[2],
			cpf[3], cpf[4], cpf[5],
			cpf[6], cpf[7], cpf[8],
			cpf[9], cpf[10])
	}

	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d",
		cpf[0], cpf[1], cpf[2], cpf[3], cpf[4],
		cpf[5], cpf[6], cpf[7], cpf[8], cpf[9], cpf[10])
}

func calculateCPFDigit(digits []int) int {
	sum := 0
	length := len(digits)

	for i, digit := range digits {
		sum += digit * (length + 1 - i)
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
