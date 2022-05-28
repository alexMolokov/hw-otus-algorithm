package simplesort

import (
	"fmt"
	"strconv"
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
	n, _ := strconv.Atoi(in[0])
	//result := Iterate(int64(n))
	//return fmt.Sprintf("%d", result)
}

func NewBubbleRunner(arr *[]int) *BubbleRunner {
	return &BubbleRunner{
		bubble: NewBubble(arr),
	}
}
