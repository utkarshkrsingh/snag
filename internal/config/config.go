package config

import (
	"fmt"

	"github.com/google/shlex"
	"go.yaml.in/yaml/v3"
)

type Config struct {
	Global GlobalConfig `yaml:"global"`
	Watch  []Rule       `yaml:"watch"`
}

type GlobalConfig struct {
	Debounce string   `yaml:"debounce"`
	Ignore   []string `yaml:"ignore"`
}

type Rule struct {
	Name     string            `yaml:"name"`
	Patterns []string          `yaml:"patterns"`
	Run      Command           `yaml:"run"`
	Restart  bool              `yaml:"restart"`
	Env      map[string]string `yaml:"env"`
}

type Command []string

func (c *Command) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		args, err := shlex.Split(value.Value)
		if err != nil {
			return err
		}
		*c = args

	case yaml.SequenceNode:
		var arr []string
		if err := value.Decode(&arr); err != nil {
			return err
		}
		*c = arr

	default:
		return fmt.Errorf("invalid type for run: expected string or list")
	}

	return nil
}
