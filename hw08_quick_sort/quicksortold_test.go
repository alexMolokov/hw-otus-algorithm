package sort

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestQuickSorOldDigits(t *testing.T) {
	t.Helper()
	taskRunner := NewQuickSortOldRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/digits/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestQuickSortOldRandom(t *testing.T) {
	t.Helper()
	taskRunner := NewQuickSortOldRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/random/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestQuickSortOldSorted(t *testing.T) {
	t.Helper()
	taskRunner := NewQuickSortOldRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/sorted/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestQuickSortOldRevers(t *testing.T) {
	t.Helper()
	taskRunner := NewQuickSortOldRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/revers/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
