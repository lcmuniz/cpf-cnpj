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
		fmt.Fprintln(os.Stderr, "Generates a random valid Brazilian CPF number and copies it to the clipboard.")
		fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	masked := flag.Bool("mask", false, "Generate CPF with mask (XXX.XXX.XXX-XX)")
	flag.Parse()

	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Failed to initialize clipboard: %v", err)
	}

	cpf := generator.GenerateCPF(*masked)

	fmt.Println(cpf)

	clipboard.Write(clipboard.FmtText, []byte(cpf))
}
