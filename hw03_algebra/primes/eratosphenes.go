package primes

import (
	"fmt"
	"math"
	"strconv"
)

func Eratosthenes(n int) int {
	count := 0
	last := int(math.Sqrt(float64(n)))
	a := make([]bool, n+1)
	for i := 2; i < len(a); i++ {
		if !a[i] {
			if i <= last {
				for j := i * i; j < len(a); j += i {
					a[j] = true
				}
			}
			count++
		}
	}
	return count
}

type EratosthenesRunner struct {
}

func (r *EratosthenesRunner) Run(in []string) string {
	n, _ := strconv.Atoi(in[0])
	result := Eratosthenes(n)
	return fmt.Sprintf("%d", result)
}

func NewEratosthenesRunner() *EratosthenesRunner {
	return &EratosthenesRunner{}
}
