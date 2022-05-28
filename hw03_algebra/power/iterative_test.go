package power

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestIterative(t *testing.T) {
	t.Helper()
	taskRunner := NewIterativeRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/../test-files/power/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
