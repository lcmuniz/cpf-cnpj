// Command cnpj prints a random valid Brazilian CNPJ and copies it to the clipboard.
package main

import (
	"github.com/lcmuniz/cpf-cnpj/internal/cli"
	"github.com/lcmuniz/cpf-cnpj/internal/generator"
)

func main() {
	cli.Tool{
		Name:      "CNPJ",
		MaskHint:  "XX.XXX.XXX/XXXX-XX",
		AlphaHelp: "generate an alphanumeric CNPJ (format issued from July 2026)",
		Generate: func(o cli.Options) string {
			if o.Alpha {
				return generator.GenerateAlphanumericCNPJ(o.Masked)
			}
			return generator.GenerateCNPJ(o.Masked)
		},
	}.Run()
}
