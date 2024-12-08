package xargs

import (
	"os"
	"testing"
)

func TestNew_NoCommandProvided(t *testing.T) {
	os.Args = []string{"./myXargs"} // Нет команды

	xg, err := New()
	if err == nil {
		t.Fatalf("expected error when no command is provided, got nil")
	}

	if xg != nil {
		t.Fatalf("expected nil XArgs object, got %v", xg)
	}
}

func TestNew_WithStdinInput(t *testing.T) {
	os.Args = []string{"./myXargs", "echo"}
	input := "arg1\narg2\narg3"

	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	w.Write([]byte(input))
	w.Close()

	os.Stdin = r

	xg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedCmd := "echo"
	if xg.cmd != expectedCmd {
		t.Fatalf("expected cmd to be %q, got %q", expectedCmd, xg.cmd)
	}

	expectedArgs := []string{"arg1", "arg2", "arg3"}
	if !equalStringSlices(xg.args, expectedArgs) {
		t.Fatalf("expected args %v, got %v", expectedArgs, xg.args)
	}
}

func TestExecute_CommandSuccess(t *testing.T) {
	os.Args = []string{"./myXargs", "echo"}
	input := "hello\nworld"

	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	w.Write([]byte(input))
	w.Close()
	os.Stdin = r

	xg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output, err := xg.Execute()
	if err != nil {
		t.Fatalf("unexpected error during execution: %v", err)
	}

	expectedOutput := "hello world\n"
	if string(output) != expectedOutput {
		t.Fatalf("expected output %q, got %q", expectedOutput, string(output))
	}
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
