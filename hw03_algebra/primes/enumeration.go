package primes

import (
	"fmt"
	"strconv"
)

func IsPrime(n int) bool {
	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	last := n / 2
	for i := 3; i <= last; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func CountPrime(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			count++
		}
	}
	return count
}

type EnumerationRunner struct {
}

func (r *EnumerationRunner) Run(in []string) string {
	n, _ := strconv.Atoi(in[0])
	result := CountPrime(n)
	return fmt.Sprintf("%d", result)
}

func NewEnumerationRunner() *EnumerationRunner {
	return &EnumerationRunner{}
}
