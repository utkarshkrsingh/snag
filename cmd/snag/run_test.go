package main

import (
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "run without args",
			args:     []string{"run"},
			expected: "runs the whole script once",
		},
		{
			name: "run with args",
			args: []string{"run", "build"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := newApplication()

			out, err := executeCommand(app.rootCmd, tt.args...)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !strings.Contains(out, tt.expected) {
				t.Fatalf("unexpected output, got: %v expected: %v", out, tt.expected)
			}
		})
	}
}
