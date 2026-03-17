package config

func (c *Config) ApplyDefaults() {
	if c.Global.Debounce == "" {
		c.Global.Debounce = "500ms"
	}
}
