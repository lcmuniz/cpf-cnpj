// Package cli holds the command-line logic shared by the cpf and cnpj tools.
package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"golang.design/x/clipboard"
)

// Tool describes one generator command.
type Tool struct {
	Name     string            // document name shown in help, e.g. "CPF"
	MaskHint string            // mask format shown in help, e.g. "XXX.XXX.XXX-XX"
	Generate func(bool) string // returns a number, masked if the argument is true
}

// Main parses the command line, prints a generated number and copies it to
// the clipboard. It returns the process exit code.
func (t Tool) Main(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {
		fmt.Fprintf(stderr, "Generates a random valid Brazilian %s number and copies it to the clipboard.\n", t.Name)
		fmt.Fprintf(stderr, "\nUsage of %s:\n", args[0])
		fs.PrintDefaults()
	}
	masked := fs.Bool("mask", false, fmt.Sprintf("format the %s as %s", t.Name, t.MaskHint))
	noCopy := fs.Bool("no-copy", false, "only print, do not copy to the clipboard")
	if err := fs.Parse(args[1:]); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}

	number := t.Generate(*masked)
	fmt.Fprintln(stdout, number)

	if !*noCopy {
		if err := copyToClipboard(number); err != nil {
			// The number was already printed, so this is only a warning.
			fmt.Fprintf(stderr, "warning: could not copy to clipboard: %v\n", err)
		}
	}
	return 0
}

// Run is a convenience wrapper for main functions.
func (t Tool) Run() {
	os.Exit(t.Main(os.Args, os.Stdout, os.Stderr))
}

func copyToClipboard(s string) error {
	if err := clipboard.Init(); err != nil {
		return err
	}
	_, err := clipboard.Write(context.Background(), clipboard.FmtText, []byte(s))
	return err
}
