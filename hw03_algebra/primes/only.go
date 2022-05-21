package primes

import (
	"fmt"
	"math"
	"strconv"
)

type Only struct {
	primes []int
}

func (p *Only) isPrime(n int) bool {
	last := int(math.Sqrt(float64(n)))
	for i := 0; p.primes[i] <= last; i++ {
		if n%p.primes[i] == 0 {
			return false
		}
	}
	p.primes = append(p.primes, n)
	return true
}

func (p *Only) CountPrime(n int) int {
	if n < 2 {
		return 0
	}

	p.primes = make([]int, 1)
	p.primes[0] = 2

	count := 0
	for i := 2; i <= n; i++ {
		if p.isPrime(i) {
			count++
		}
	}
	return count
}

type OnlyRunner struct {
}

func (r *OnlyRunner) Run(in []string) string {
	n, _ := strconv.Atoi(in[0])
	only := Only{}
	result := only.CountPrime(n)
	return fmt.Sprintf("%d", result)
}

func NewOnlyRunner() *OnlyRunner {
	return &OnlyRunner{}
}
