package fibo

import (
	"fmt"
	"strconv"
)

func Recursion(n int) uint {
	if n == 0 || n == 1 {
		return uint(n)
	}
	return Recursion(n-2) + Recursion(n-1)
}

type RecursionRunner struct {
}

func (r *RecursionRunner) Run(in []string) string {
	n, _ := strconv.Atoi(in[0])
	result := Recursion(n)
	return fmt.Sprintf("%d", result)
}

func NewRecursionRunner() *RecursionRunner {
	return &RecursionRunner{}
}
