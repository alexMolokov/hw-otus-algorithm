package heapsort

import (
	"fmt"
	"strconv"
	"strings"
)

type Selection struct{}

func (s *Selection) Sort(arr *[]int) []int {
	size := len(*arr)
	for i := size; i > 0; i-- {
		index := s.max(arr, i)
		s.swap(arr, i-1, index)
	}

	return *arr
}

func (s *Selection) swap(arr *[]int, i, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func (s *Selection) max(arr *[]int, size int) int {
	var max int
	for i := 0; i < size; i++ {
		if (*arr)[i] > (*arr)[max] {
			max = i
		}
	}
	return max
}

func NewSelectionSort() *Selection {
	return &Selection{}
}

type SelectionRunner struct {
	selection *Selection
}

func (s *SelectionRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	result := fmt.Sprintf("%v", s.selection.Sort(&arr))
	return result[1 : len(result)-1]
}

func NewSelectionRunner(selection *Selection) *SelectionRunner {
	return &SelectionRunner{
		selection: selection,
	}
}
