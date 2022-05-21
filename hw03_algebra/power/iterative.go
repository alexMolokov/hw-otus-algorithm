package power

import (
	"fmt"
	"strconv"
)

// Iterative возведение числа a в положительную степень n
// через последовательное умножение
func Iterative(a float64, n int) float64 {
	var result float64 = 1
	for i := 1; i <= n; i++ {
		result *= a
	}

	return result
}

type IterativeRunner struct {
}

func (r *IterativeRunner) Run(in []string) string {
	a, _ := strconv.ParseFloat(in[0], 64)
	n, _ := strconv.Atoi(in[1])
	result := Iterative(a, n)
	return fmt.Sprintf("%.11f", result)
}

func NewIterativeRunner() *IterativeRunner {
	return &IterativeRunner{}
}
