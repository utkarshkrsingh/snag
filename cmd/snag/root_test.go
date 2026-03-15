package main

import (
	"strings"
	"testing"
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "without verbose",
			args:     []string{},
			expected: "started watching",
		},
		{
			name:     "with verbose",
			args:     []string{"-v"},
			expected: "verbose mode enabled",
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
				t.Fatalf("unexpected ouput, got %v expected: %v", out, tt.expected)
			}
		})
	}
}
