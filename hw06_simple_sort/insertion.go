package simplesort

import (
	"fmt"
	"strconv"
	"strings"
)

type Insertion struct {
	arr []int
}

func (s *Insertion) swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s *Insertion) Sort() []int {
	size := len(s.arr)
	for i := 1; i < size; i++ {
		for j := i; j > 0; j-- {
			if s.arr[j] > s.arr[j-1] {
				break
			}
			s.swap(j, j-1)
		}
	}
	return s.arr
}

func NewInsertion(arr *[]int) *Insertion {
	return &Insertion{
		arr: *arr,
	}
}

type InsertionRunner struct {
	insertion *Insertion
}

func (r *InsertionRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	r.insertion = NewInsertion(&arr)

	result := fmt.Sprintf("%v", r.insertion.Sort())
	return result[1 : len(result)-1]
}

func NewInsertionRunner() *InsertionRunner {
	return &InsertionRunner{}
}
