package sort

import (
	"fmt"
	"strconv"
	"strings"
)

type QuickSortOld struct {
	arr []int
}

func (s *QuickSortOld) Sort() []int {
	s.sort(0, len(s.arr)-1)
	return s.arr
}

func (s *QuickSortOld) sort(l, r int) {
	if r <= l {
		return
	}
	m := s.split(l, r)
	s.sort(l, m-1)
	s.sort(m+1, r)
}

func (s *QuickSortOld) split(l, r int) int {
	pivot := s.arr[r]
	i, j, k := l, r-1, 0

	for i < j {
		if k == 0 {
			if s.arr[i] <= pivot {
				i++
				continue
			}
			k++
		}

		if k == 1 {
			if s.arr[j] > pivot {
				j--
				continue
			}
			k++
		}

		if k == 2 {
			s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
			k = 0
		}
	}

	if s.arr[i] < s.arr[r] {
		i++
	}

	s.arr[i], s.arr[r] = s.arr[r], s.arr[i]

	return i
}

func NewQuickSortOld(arr *[]int) *QuickSortOld {
	return &QuickSortOld{
		arr: *arr,
	}
}

type QuickSortOldRunner struct {
	quick *QuickSortOld
}

func (r *QuickSortOldRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	r.quick = NewQuickSortOld(&arr)

	result := fmt.Sprintf("%v", r.quick.Sort())
	return result[1 : len(result)-1]
}

func NewQuickSortOldRunner() *QuickSortOldRunner {
	return &QuickSortOldRunner{}
}
