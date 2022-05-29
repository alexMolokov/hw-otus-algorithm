package simplesort

import (
	"fmt"
	"strconv"
	"strings"
)

type InsertionBinary struct {
	Insertion
}

func (s *InsertionBinary) BinaryPosition(left, right, value int) int {
	m := (left + right) / 2

	for {
		if s.arr[m] <= value {
			left = m + 1
		}

		if s.arr[m] > value {
			right = m - 1
		}

		if left >= right {
			if s.arr[left] < value {
				return left + 1
			}
			return left
		}

		m = (left + right) / 2
	}
}

func (s *InsertionBinary) Sort() []int {
	size := len(s.arr)
	for i := 1; i < size; i++ {
		k := s.BinaryPosition(0, i-1, s.arr[i])

		for j := i - 1; j >= k; j-- {
			s.swap(j, j+1)
		}
	}
	return s.arr
}

func NewInsertionBinary(arr *[]int) *InsertionBinary {
	return &InsertionBinary{
		Insertion{arr: *arr},
	}
}

type InsertionBinaryRunner struct {
	insertion *InsertionBinary
}

func (r *InsertionBinaryRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	r.insertion = NewInsertionBinary(&arr)

	result := fmt.Sprintf("%v", r.insertion.Sort())
	return result[1 : len(result)-1]
}

func NewInsertionBinaryRunner() *InsertionBinaryRunner {
	return &InsertionBinaryRunner{}
}
