package task

type TaskModel struct {
	Name    string
	Command string
	Args    []string
	Status  TaskStatus
}
