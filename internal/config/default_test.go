package config

import "testing"

func TestCondig_ApplyDefaults(t *testing.T) {
	tests := []struct {
		name         string
		initialCfg   Config
		wantDebounce string
	}{
		{
			name:         "Empty debounce gets default",
			initialCfg:   Config{Global: GlobalConfig{Debounce: ""}},
			wantDebounce: "500ms",
		},
		{
			name:         "Existing debounce is respected",
			initialCfg:   Config{Global: GlobalConfig{Debounce: "2s"}},
			wantDebounce: "2s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initialCfg.ApplyDefaults()

			if tt.initialCfg.Global.Debounce != tt.wantDebounce {
				t.Errorf("ApplyDefaults() debounce = %v, want %v", tt.initialCfg.Global.Debounce, tt.wantDebounce)
			}
		})
	}
}
