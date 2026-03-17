package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"go.yaml.in/yaml/v3"
)

func TestConfigMgr_Load(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name        string
		yamlContent string
		wantErr     bool
		errContains string
	}{
		{
			name: "Valid configuration",
			yamlContent: `
global:
  debounce: 500ms
  ignore:
    - .git
watch:
  - name: server
    patterns: ["**/*.go"]
    run: go run main.go
`,
			wantErr: false,
		},
		{
			name: "Unknown field (Struct decoding check)",
			yamlContent: `
lobal:  # <--- Typo here
  debounce: 500ms
`,
			wantErr:     true,
			errContains: "field lobal not found",
		},
		{
			name:    "File not found",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := filepath.Join(tempDir, tt.name+".yaml")

			if tt.yamlContent != "" {
				err := os.WriteFile(filePath, []byte(tt.yamlContent), 0644)
				if err != nil {
					t.Fatalf("Failde to write temp file: %v", err)
				}
			} else {
				filePath = filepath.Join(tempDir, "missing.yaml")
			}

			mgr := New(filePath)
			err := mgr.Load()

			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.errContains != "" {
				if err != nil && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Load() error = %v, expected it to contain %v", err, tt.errContains)
				}
			}
		})
	}
}

func TestCommand_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name        string
		yamlContent string
		want        Command
		wantErr     bool
	}{
		{
			name:        "Parse from String (Scalar)",
			yamlContent: `run: go run main.go`,
			want:        Command{"go", "run", "main.go"},
			wantErr:     false,
		},
		{
			name:        "Parse from String with Quotes (shlex)",
			yamlContent: `run: gcc "my file.c" -o app`,
			want:        Command{"gcc", "my file.c", "-o", "app"},
			wantErr:     false,
		},
		{
			name: "Parse from List (Sequence)",
			yamlContent: `
run:
  - golangci-lint
  - run
`,
			want:    Command{"golangci-lint", "run"},
			wantErr: false,
		},
		{
			name: "Invalid Type (Map/Object)",
			yamlContent: `
run:
  command: go run main.go
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dummy struct {
				Run Command `yaml:"run"`
			}

			err := yaml.Unmarshal([]byte(tt.yamlContent), &dummy)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(dummy.Run, tt.want) {
				t.Errorf("UnmarshalYAML() got: %v, expected: %v", dummy.Run, tt.want)
			}
		})
	}
}
