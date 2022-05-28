package power

import (
	"fmt"
	"strconv"
)

// BinaryExpansion возведение числа a в положительную степень n
// через двоичное разложение показателя степени
func BinaryExpansion(a float64, n int) float64 {
	if n == 0 {
		return 1
	}

	var d float64 = 1
	s := a
	for power := n; power > 1; power = power / 2 {
		if power%2 == 1 {
			d *= s
		}
		s *= s
	}
	d *= s
	return d
}

type BinaryExpansionRunner struct {
}

func (b *BinaryExpansionRunner) Run(in []string) string {
	a, _ := strconv.ParseFloat(in[0], 64)
	n, _ := strconv.Atoi(in[1])
	result := BinaryExpansion(a, n)
	return fmt.Sprintf("%.11f", result)
}

func NewBinaryExpansionRunner() *BinaryExpansionRunner {
	return &BinaryExpansionRunner{}
}
