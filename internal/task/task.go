package task

type Task struct {
	Command   string   `yaml:"cmd"`
	Args      []string `yaml:"args"`
	Parallel  []string `yaml:"parallel,omitempty"`
	DependsOn []string `yaml:"depends_on,omitempty"`
}

type Config struct {
	Tasks map[string]Task `yaml:"tasks"`
}

type TaskNode struct {
	Name      string
	Command   string
	Args      []string
	DependsOn []string
}
