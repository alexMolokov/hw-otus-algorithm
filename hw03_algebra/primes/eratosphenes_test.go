package primes

import (
	"github.com/alexMolokov/hw-otus-algorithm/system"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestEratosthenes(t *testing.T) {
	require.Equal(t, 4, Eratosthenes(10))
	t.Helper()
	taskRunner := NewEratosthenesRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/../test-files/primes/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
