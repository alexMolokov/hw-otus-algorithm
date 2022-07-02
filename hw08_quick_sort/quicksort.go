package sort

import (
	"fmt"
	"strconv"
	"strings"
)

type QuickSort struct {
	arr []int
}

func (s *QuickSort) Sort() []int {
	s.sort(0, len(s.arr)-1)
	return s.arr
}

func (s *QuickSort) sort(l, r int) {
	if r <= l {
		return
	}
	m := s.split(l, r)
	s.sort(l, m-1)
	s.sort(m+1, r)
}

func (s *QuickSort) split(l, r int) int {
	pivot := s.arr[r]
	m := l - 1
	for i := l; i <= r; i++ {
		if s.arr[i] <= pivot {
			m++
			s.arr[m], s.arr[i] = s.arr[i], s.arr[m]
		}
	}
	return m
}

func NewQuickSort(arr *[]int) *QuickSort {
	return &QuickSort{
		arr: *arr,
	}
}

type QuickSortRunner struct {
	quick *QuickSort
}

func (r *QuickSortRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	r.quick = NewQuickSort(&arr)

	result := fmt.Sprintf("%v", r.quick.Sort())
	return result[1 : len(result)-1]
}

func NewQuickSortRunner() *QuickSortRunner {
	return &QuickSortRunner{}
}
