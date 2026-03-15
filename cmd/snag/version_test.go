package main

import (
	"testing"
)

func TestVersionCmd(t *testing.T) {
	app := newApplication()

	out, err := executeCommand(app.rootCmd, "version")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(out) == 0 {
		t.Fatalf("unexpected output")
	}
}
