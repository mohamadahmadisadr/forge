package task

type TaskStatus string

const (
	Done    TaskStatus = "done"
	Running TaskStatus = "running"
	None    TaskStatus = "none"
	Failed  TaskStatus = "failed"
)
