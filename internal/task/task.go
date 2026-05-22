package task

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Task struct {
	Command  string   `yaml:"cmd"`
	Args     []string `yaml:"args"`
	Parallel []string `yaml:"parallel,omitempty"`
}

type YamlConfig struct {
	Tasks map[string]Task `yaml:"tasks"`
}

func GetYamlConfig() (*YamlConfig, error) {

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config YamlConfig

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
