package cli

import (
	"bytes"
	"strings"
	"testing"
)

func fakeTool() Tool {
	return Tool{
		Name:     "TEST",
		MaskHint: "X-X",
		Generate: func(masked bool) string {
			if masked {
				return "1-2"
			}
			return "12"
		},
	}
}

func TestMainPrintsNumber(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"test", "-no-copy"}, "12\n"},
		{[]string{"test", "-no-copy", "-mask"}, "1-2\n"},
	}
	for _, tt := range tests {
		var out, errOut bytes.Buffer
		if code := fakeTool().Main(tt.args, &out, &errOut); code != 0 {
			t.Fatalf("Main(%v) exit code = %d, stderr: %s", tt.args, code, errOut.String())
		}
		if out.String() != tt.want {
			t.Errorf("Main(%v) printed %q, want %q", tt.args, out.String(), tt.want)
		}
	}
}

func TestMainHelp(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := fakeTool().Main([]string{"test", "-h"}, &out, &errOut); code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if !strings.Contains(errOut.String(), "Brazilian TEST number") || out.Len() != 0 {
		t.Errorf("unexpected help output: stdout=%q stderr=%q", out.String(), errOut.String())
	}
}

func TestMainBadFlag(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := fakeTool().Main([]string{"test", "-nope"}, &out, &errOut); code != 2 {
		t.Fatalf("exit code = %d, want 2", code)
	}
}
