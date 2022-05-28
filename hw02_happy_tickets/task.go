package happytickets

import (
	"fmt"
	"strconv"
)

type Task struct{}

func (t *Task) Get(n int) int64 {
	if n <= 0 {
		return 0
	}

	matrix := make([][]int64, 9*n+1)
	for i := 0; i <= 9*n; i++ {
		matrix[i] = make([]int64, n)
	}

	for i := 0; i <= 9; i++ {
		matrix[i][0] = 1
	}

	current := 1
	for current < n {
		for i := 0; i <= 9*(current+1); i++ {
			for j := 0; j <= 9; j++ {
				if i < j {
					continue
				}
				matrix[i][current] += matrix[i-j][current-1]
			}
		}
		current++
	}

	var sum int64
	for i := 0; i <= 9*n; i++ {
		val := matrix[i][n-1]
		sum += val * val
	}

	return sum
}

func NewTask() *Task {
	return &Task{}
}

// TaskRunner структура для запуска в системе тестирования
// задача данной структуры получить входные данные преобразовать их в необходимый формат
// вызвать метод структуры Task, получить от нее ответ и преобразовать в необходимый формат ответа.
type TaskRunner struct {
	Task *Task
}

func (t *TaskRunner) Run(in []string) string {
	n, _ := strconv.Atoi(in[0])
	result := t.Task.Get(n)
	return fmt.Sprintf("%d", result)
}

func NewTaskRunner(task *Task) *TaskRunner {
	return &TaskRunner{
		Task: task,
	}
}
