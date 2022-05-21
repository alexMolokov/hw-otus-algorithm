package fibo

import (
	"fmt"
	"math/big"
	"strconv"
)

func Iterate(n int64) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(n)
	}
	prev, prev2 := big.NewInt(1), big.NewInt(0)
	for i := int64(2); i <= n; i++ {
		temp := prev
		prev = big.NewInt(0).Add(prev, prev2)
		prev2 = temp
	}
	return prev
}

type IterateRunner struct {
}

func (r *IterateRunner) Run(in []string) string {
	n, _ := strconv.Atoi(in[0])
	result := Iterate(int64(n))
	return fmt.Sprintf("%d", result)
}

func NewIterateRunner() *IterateRunner {
	return &IterateRunner{}
}
