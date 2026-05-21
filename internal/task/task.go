package task

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Task struct {
	Name    string   `yaml:"name"`
	Command string   `yaml:"cmd"`
	Args    []string `yaml:"args"`
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
