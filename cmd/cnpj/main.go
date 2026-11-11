// Command cnpj prints a random valid Brazilian CNPJ and copies it to the clipboard.
package main

import (
	"github.com/lcmuniz/cpf-cnpj/internal/cli"
	"github.com/lcmuniz/cpf-cnpj/internal/generator"
)

func main() {
	cli.Tool{
		Name:     "CNPJ",
		MaskHint: "XX.XXX.XXX/XXXX-XX",
		Generate: generator.GenerateCNPJ,
	}.Run()
}
