package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lcmuniz/cpf-cnpj/internal/generator"
	"golang.design/x/clipboard"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Generates a random valid Brazilian CNPJ number and copies it to the clipboard.")
		fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	masked := flag.Bool("mask", false, "Generate CNPJ with mask (XX.XXX.XXX/XXXX-XX)")
	flag.Parse()

	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Failed to initialize clipboard: %v", err)
	}

	cnpj := generator.GenerateCNPJ(*masked)

	fmt.Println(cnpj)

	clipboard.Write(clipboard.FmtText, []byte(cnpj))
}
