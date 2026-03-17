package config

import (
	"strings"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	validRule := Rule{
		Name:     "build",
		Patterns: []string{"**/*.go"},
		Run:      Command{"go", "build"},
	}

	tests := []struct {
		name        string
		cfg         Config
		wantErr     bool
		errContains string
	}{
		{
			name: "Valid Configuration",
			cfg: Config{
				Global: GlobalConfig{Debounce: "500ms"},
				Watch:  []Rule{validRule},
			},
			wantErr: false,
		},
		{
			name: "Missing Watch Rules",
			cfg: Config{
				Global: GlobalConfig{Debounce: "500ms"},
				Watch:  []Rule{},
			},
			wantErr:     true,
			errContains: "no watch rules defined",
		},
		{
			name: "Invalid Debounce Format",
			cfg: Config{
				Global: GlobalConfig{Debounce: "500"},
				Watch:  []Rule{validRule},
			},
			wantErr:     true,
			errContains: "invalid global debounce duration",
		},
		{
			name: "Rule Missing Name",
			cfg: Config{
				Global: GlobalConfig{Debounce: "500ms"},
				Watch: []Rule{{
					Name:     "   ",
					Patterns: []string{"**/*.go"},
					Run:      Command{"go", "build"},
				}},
			},
			wantErr:     true,
			errContains: "missing a 'name'",
		},
		{
			name: "Rule Missing Patterns",
			cfg: Config{
				Global: GlobalConfig{Debounce: "500ms"},
				Watch: []Rule{{
					Name:     "test",
					Patterns: []string{},
					Run:      Command{"go", "test"},
				}},
			},
			wantErr:     true,
			errContains: "must have at least one pattern",
		},
		{
			name: "Rule Missing Command",
			cfg: Config{
				Global: GlobalConfig{Debounce: "500ms"},
				Watch: []Rule{{
					Name:     "run",
					Patterns: []string{"**/*.go"},
					Run:      Command{},
				}},
			},
			wantErr:     true,
			errContains: "missing a 'run' command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.errContains != "" {
				if err != nil && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Validate() error = %v, expected it to contain %q", err, tt.errContains)
				}
			}
		})
	}
}
