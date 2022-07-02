package heapsort

import (
	"fmt"
	"strconv"
	"strings"
)

type Heap struct {
	arr []int
}

func (h *Heap) left(i int) int {
	return 2*i + 1
}

func (h *Heap) right(i int) int {
	return 2*i + 2
}

func (h *Heap) makeHeap() {
	size := len(h.arr)
	for i := size/2 - 1; i >= 0; i-- {
		h.heapify(i, size)
	}
}

func (h *Heap) Sort() []int {
	size := len(h.arr)
	for i := size - 1; i > 0; i-- {
		h.swap(0, i)
		h.heapify(0, i)
	}
	return h.arr
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) heapify(root int, size int) {
	left := h.left(root)
	right := h.right(root)

	max := root
	if left < size && h.arr[left] > h.arr[max] {
		max = left
	}

	if right < size && h.arr[right] > h.arr[max] {
		max = right
	}

	if max == root {
		return
	}

	h.swap(max, root)
	h.heapify(max, size)
}

func NewHeap(in *[]int) *Heap {
	h := &Heap{
		arr: *in,
	}
	h.makeHeap()
	return h
}

type HeapRunner struct {
	heap *Heap
}

func (h *HeapRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	h.heap = NewHeap(&arr)

	result := fmt.Sprintf("%v", h.heap.Sort())
	return result[1 : len(result)-1]
}

func NewHeapRunner() *HeapRunner {
	return &HeapRunner{}
}
