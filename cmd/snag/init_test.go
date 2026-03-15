package main

import (
	"strings"
	"testing"
)

func TestInitCmd(t *testing.T) {
	app := newApplication()

	out, err := executeCommand(app.rootCmd, "init")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, "config initiated") {
		t.Fatalf("unexpected output: %v", out)
	}

}
