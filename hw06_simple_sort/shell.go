package simplesort

import (
	"fmt"
	"strconv"
	"strings"
)

type Shell struct {
	arr []int
}

func (sh *Shell) swap(i, j int) {
	sh.arr[i], sh.arr[j] = sh.arr[j], sh.arr[i]
}

func (sh *Shell) Sort() []int {
	size := len(sh.arr)
	for s := size / 2; s > 0; s /= 2 {
		for i := s; i < size; i++ {
			for j := i - s; j >= 0 && sh.arr[j] > sh.arr[j+s]; j -= s {
				sh.swap(j, j+s)
			}
		}
	}
	return sh.arr
}

func NewShell(arr *[]int) *Shell {
	return &Shell{
		arr: *arr,
	}
}

type ShellRunner struct {
	Shell *Shell
}

func (r *ShellRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	r.Shell = NewShell(&arr)

	result := fmt.Sprintf("%v", r.Shell.Sort())
	return result[1 : len(result)-1]
}

func NewShellRunner() *ShellRunner {
	return &ShellRunner{}
}
