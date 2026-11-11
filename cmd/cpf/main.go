// Command cpf prints a random valid Brazilian CPF and copies it to the clipboard.
package main

import (
	"github.com/lcmuniz/cpf-cnpj/internal/cli"
	"github.com/lcmuniz/cpf-cnpj/internal/generator"
)

func main() {
	cli.Tool{
		Name:     "CPF",
		MaskHint: "XXX.XXX.XXX-XX",
		Generate: func(o cli.Options) string { return generator.GenerateCPF(o.Masked) },
	}.Run()
}
