package primes

import (
	"github.com/alexMolokov/hw-otus-algorithm/system"
	"os"
	"path/filepath"
	"testing"
)

func TestP(t *testing.T) {
	t.Helper()
	taskRunner := NewEnumerationRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/../test-files/primes/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
