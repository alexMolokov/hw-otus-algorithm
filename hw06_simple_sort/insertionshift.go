package simplesort

import (
	"fmt"
	"strconv"
	"strings"
)

type InsertionShift struct {
	Insertion
}

func (s *InsertionShift) Sort() []int {
	size := len(s.arr)
	for i := 1; i < size; i++ {
		el := s.arr[i]
		for j := i - 1; j >= 0; j-- {
			if s.arr[j] <= el {
				s.arr[j+1] = el
				break
			}
			s.swap(j, j+1)
		}
	}
	return s.arr
}

func NewInsertionShift(arr *[]int) *InsertionShift {
	return &InsertionShift{
		Insertion{arr: *arr},
	}
}

type InsertionShiftRunner struct {
	insertion *InsertionShift
}

func (r *InsertionShiftRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	r.insertion = NewInsertionShift(&arr)

	result := fmt.Sprintf("%v", r.insertion.Sort())
	return result[1 : len(result)-1]
}

func NewInsertionShiftRunner() *InsertionShiftRunner {
	return &InsertionShiftRunner{}
}
