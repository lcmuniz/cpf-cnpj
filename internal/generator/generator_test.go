package generator

import (
	"regexp"
	"testing"
)

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"52998224725", true},
		{"529.982.247-25", true},
		{"12345678909", true},
		{"52998224724", false},  // wrong second check digit
		{"52998224715", false},  // wrong first check digit
		{"11111111111", false},  // repeated digits
		{"5299822472", false},   // too short
		{"529982247250", false}, // too long
		{"529.982.247-2a", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := IsValidCPF(tt.in); got != tt.want {
			t.Errorf("IsValidCPF(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestIsValidCNPJ(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"11222333000181", true},
		{"11.222.333/0001-81", true},
		{"11222333000180", false}, // wrong second check digit
		{"11222333000191", false}, // wrong first check digit
		{"00000000000000", false}, // repeated digits
		{"1122233300018", false},  // too short
		{"11.222.333/0001-8x", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := IsValidCNPJ(tt.in); got != tt.want {
			t.Errorf("IsValidCNPJ(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestGenerateCPF(t *testing.T) {
	plain := regexp.MustCompile(`^\d{11}$`)
	masked := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	for range 10000 {
		if c := GenerateCPF(false); !plain.MatchString(c) || !IsValidCPF(c) {
			t.Fatalf("GenerateCPF(false) = %q, not a valid plain CPF", c)
		}
		if c := GenerateCPF(true); !masked.MatchString(c) || !IsValidCPF(c) {
			t.Fatalf("GenerateCPF(true) = %q, not a valid masked CPF", c)
		}
	}
}

func TestGenerateCNPJ(t *testing.T) {
	plain := regexp.MustCompile(`^\d{14}$`)
	masked := regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)
	for range 10000 {
		if c := GenerateCNPJ(false); !plain.MatchString(c) || !IsValidCNPJ(c) {
			t.Fatalf("GenerateCNPJ(false) = %q, not a valid plain CNPJ", c)
		}
		if c := GenerateCNPJ(true); !masked.MatchString(c) || !IsValidCNPJ(c) {
			t.Fatalf("GenerateCNPJ(true) = %q, not a valid masked CNPJ", c)
		}
	}
}

func TestRandomCharsNeverAllEqual(t *testing.T) {
	for range 100000 {
		if d := randomChars(2, digitChars); allEqual(d) {
			t.Fatalf("randomChars returned repeated characters %q", d)
		}
	}
}

func TestIsValidCNPJAlphanumeric(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Worked example from the Receita Federal Q&A on the alphanumeric CNPJ.
		{"12ABC34501DE35", true},
		{"12.ABC.345/01DE-35", true},
		{"12ABC34501DE36", false}, // wrong second check digit
		{"12ABC34501DE45", false}, // wrong first check digit
		{"12abc34501de35", false}, // letters must be uppercase
		{"12ABC34501DE3A", false}, // check digits must be numeric
		{"12ABC34501D-35", false}, // too short
		{"AAAAAAAAAAAAAA", false},
		{"12.ABC.345/01DE-3$", false},
	}
	for _, tt := range tests {
		if got := IsValidCNPJ(tt.in); got != tt.want {
			t.Errorf("IsValidCNPJ(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestGenerateAlphanumericCNPJ(t *testing.T) {
	plain := regexp.MustCompile(`^[0-9A-Z]{12}\d{2}$`)
	masked := regexp.MustCompile(`^[0-9A-Z]{2}\.[0-9A-Z]{3}\.[0-9A-Z]{3}/[0-9A-Z]{4}-\d{2}$`)
	letter := regexp.MustCompile(`[A-Z]`)
	for range 10000 {
		c := GenerateAlphanumericCNPJ(false)
		if !plain.MatchString(c) || !letter.MatchString(c) || !IsValidCNPJ(c) {
			t.Fatalf("GenerateAlphanumericCNPJ(false) = %q, not a valid alphanumeric CNPJ", c)
		}
		if c := GenerateAlphanumericCNPJ(true); !masked.MatchString(c) || !IsValidCNPJ(c) {
			t.Fatalf("GenerateAlphanumericCNPJ(true) = %q, not a valid masked CNPJ", c)
		}
	}
}
