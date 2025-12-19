package config

type ConfigManager interface {
	LoadConfig()
}

type Config struct{}

func NewConfigManager() *Config {
	return &Config{}
}

func (cm *Config) LoadConfig() {}
