package heapsort

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestSelectionSortDigits(t *testing.T) {
	t.Helper()
	taskRunner := NewSelectionRunner(NewSelectionSort())
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/digits/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestSelectionSortSorted(t *testing.T) {
	t.Helper()
	taskRunner := NewSelectionRunner(NewSelectionSort())
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/sorted/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestSelectionSortRandom(t *testing.T) {
	t.Helper()
	taskRunner := NewSelectionRunner(NewSelectionSort())
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/random/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestSelectionSortRevers(t *testing.T) {
	t.Helper()
	taskRunner := NewSelectionRunner(NewSelectionSort())
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/revers/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
