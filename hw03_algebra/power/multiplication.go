package power

import (
	"fmt"
	"strconv"
)

// Multiplication возведение числа a в положительную степень n
// через домножение O(N/2+LogN) = O(N)
func Multiplication(a float64, n int) float64 {
	if n == 0 {
		return 1
	}

	result := a
	s := 2

	for s < n {
		result *= result
		s *= 2
	}

	if n > s/2 {
		result = result * Iterative(a, n-s/2)
	}

	fmt.Printf("степень 2 для %d %d", n, s/2)

	return result
}

type MultiplicationRunner struct {
}

func (r *MultiplicationRunner) Run(in []string) string {
	a, _ := strconv.ParseFloat(in[0], 64)
	n, _ := strconv.Atoi(in[1])
	result := Multiplication(a, n)
	return fmt.Sprintf("%.11f", result)
}

func NewMultiplicationRunner() *MultiplicationRunner {
	return &MultiplicationRunner{}
}
