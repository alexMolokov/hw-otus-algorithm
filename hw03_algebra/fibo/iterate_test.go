package fibo

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestIterate(t *testing.T) {
	t.Helper()
	taskRunner := NewIterateRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/../test-files/fibo/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
