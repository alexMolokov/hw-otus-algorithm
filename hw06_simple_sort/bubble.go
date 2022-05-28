package simplesort

import (
	"fmt"
	"strconv"
	"strings"
)

type Bubble struct {
	arr []int
}

func (b *Bubble) Sort() []int {
	size := len(b.arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if b.arr[j] > b.arr[j+1] {
				b.swap(j, j+1)
			}
		}
	}
	return b.arr
}

func (b *Bubble) swap(i, j int) {
	b.arr[i], b.arr[j] = b.arr[j], b.arr[i]
}

func NewBubble(arr *[]int) *Bubble {
	return &Bubble{
		arr: *arr,
	}
}

type BubbleRunner struct {
	bubble *Bubble
}

func (r *BubbleRunner) Run(in []string) string {
	size, _ := strconv.Atoi(in[0])
	data := strings.Split(in[1], " ")

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(data[i])
	}

	r.bubble = NewBubble(&arr)

	result := fmt.Sprintf("%v", r.bubble.Sort())
	return result[1 : len(result)-1]
}

func NewBubbleRunner(arr *[]int) *BubbleRunner {
	return &BubbleRunner{
		bubble: NewBubble(arr),
	}
}
