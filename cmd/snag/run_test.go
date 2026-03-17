package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/utkarshkrsingh/snag/internal/config"
	"github.com/utkarshkrsingh/snag/internal/ui"
)

func TestRunCmd(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "snag.yaml")

	dummyConfig := `
global:
    debounce: 500ms
watch:
    - name: test_task
      patterns: ["**/*.go"]
      run: echo "hello test"
`

	err := os.WriteFile(configPath, []byte(dummyConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to write temp config: %v", err)
	}

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "run_without_args",
			args: []string{},
		},
		{
			name: "run_with_args",
			args: []string{"test_task"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := newApplication()
			app.configMgr = *config.New(configPath)

			app.ui = ui.New(io.Discard)

			cmd := newRunCmd(app)
			cmd.SetArgs(tt.args)

			cmd.SetOut(new(bytes.Buffer))
			cmd.SetErr(new(bytes.Buffer))

			err := cmd.Execute()

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
