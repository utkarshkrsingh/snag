package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const configFileName = "snag.yaml"

type CfgMgr interface {
	LoadConfig(dir string) error
	GenerateConfig(dir string) error
}

type ConfigManager struct {
	Configs *Config
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}

func (cm *ConfigManager) LoadConfig(dir string) error {
	v := viper.New()
	v.SetConfigFile(filepath.Join(dir, configFileName))

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return err
	}

	cm.Configs = &config
	return nil
}

func (cm *ConfigManager) GenerateConfig(dir string) error {
	_, err := os.Stat(filepath.Join(dir, configFileName))
	if err == nil {
		return nil
	}

	data := []byte(template)
	err = os.WriteFile(configFileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
