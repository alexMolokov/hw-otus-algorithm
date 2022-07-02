package sort

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestQuickSortDigits(t *testing.T) {
	t.Helper()
	taskRunner := NewQuickSortRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/digits/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestQuickSortRandom(t *testing.T) {
	t.Helper()
	taskRunner := NewQuickSortRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/random/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestQuickSortSorted(t *testing.T) {
	t.Helper()
	taskRunner := NewQuickSortRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/sorted/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestQuickSortRevers(t *testing.T) {
	t.Helper()
	taskRunner := NewQuickSortRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/revers/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
