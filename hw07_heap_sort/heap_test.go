package heapsort

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestHeapSortDigits(t *testing.T) {
	t.Helper()
	taskRunner := NewHeapRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/digits/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestHeapSortSorted(t *testing.T) {
	t.Helper()
	taskRunner := NewHeapRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/sorted/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestHeapSortRandom(t *testing.T) {
	t.Helper()
	taskRunner := NewHeapRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/random/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestHeapSortRevers(t *testing.T) {
	t.Helper()
	taskRunner := NewHeapRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/revers/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
