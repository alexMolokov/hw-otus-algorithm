package sort

import (
	"fmt"
	"strconv"
	"strings"
)

type MergeSort struct {
	arr []int
}

func (s *MergeSort) Sort() []int {
	s.sort(0, len(s.arr)-1)
	return s.arr
}

func (s *MergeSort) sort(l, r int) {
	if l == r {
		return
	}

	m := (l + r) / 2

	s.sort(l, m)
	s.sort(m+1, r)
	s.merge(l, r, m)
}

func (s *MergeSort) merge(l, r, m int) {
	a := make([]int, r-l+1)

	i, j, k := l, m+1, 0
	for i <= m && j <= r {
		if s.arr[i] <= s.arr[j] {
			a[k] = s.arr[i]
			i++
		} else {
			a[k] = s.arr[j]
			j++
		}
		k++
	}

	for d := i; d <= m; d++ {
		a[k] = s.arr[d]
		k++
	}

	for d := j; d <= r; d++ {
		a[k] = s.arr[d]
		k++
	}

	for i := 0; i < len(a); i++ {
		s.arr[l+i] = a[i]
	}
}

func NewMergeSort(arr *[]int) *MergeSort {
	return &MergeSort{
		arr: *arr,
	}
}

type MergeSortRunner struct {
	quick *MergeSort
}

func (r *MergeSortRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	r.quick = NewMergeSort(&arr)

	result := fmt.Sprintf("%v", r.quick.Sort())
	return result[1 : len(result)-1]
}

func NewMergeSortRunner() *MergeSortRunner {
	return &MergeSortRunner{}
}
