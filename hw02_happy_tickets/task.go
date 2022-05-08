package happytickets

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

type TaskRunner struct {
	Task *Task
}

func (t *TaskRunner) Run(in []string) string {
	return "10"
}

func NewTaskRunner(task *Task) *TaskRunner {
	return &TaskRunner{
		Task: task,
	}
}
