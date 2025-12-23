package config

type Config struct {
	Debounce  string            `yaml:"debounce"`
	Env       map[string]string `yaml:"env"`
	Ignore    []string          `yaml:"ignore"`
	Pipelines []Pipelines       `yaml:"pipelines"`
}

type Pipelines struct {
	Name       string            `yaml:"name"`
	Extensions []string          `yaml:"extensions"`
	Command    string            `yaml:"command"`
	Env        map[string]string `yaml:"env"`
}
