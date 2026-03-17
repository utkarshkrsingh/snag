package config

import (
	"bytes"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type ConfigMgr struct {
	Config   Config
	FilePath string
}

func New(filePath string) *ConfigMgr {
	return &ConfigMgr{
		FilePath: filePath,
	}
}

func (cfgMgr *ConfigMgr) Load() error {
	data, err := os.ReadFile(cfgMgr.FilePath)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)

	decoder.KnownFields(true)

	if err := decoder.Decode(&cfgMgr.Config); err != nil {
		return fmt.Errorf("invalid yaml format: %w", err)
	}

	return nil
}
