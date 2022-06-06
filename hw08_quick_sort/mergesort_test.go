package sort

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestMergeSortDigits(t *testing.T) {
	t.Helper()
	taskRunner := NewMergeSortRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/digits/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestMergeSortRandom(t *testing.T) {
	t.Helper()
	taskRunner := NewMergeSortRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/random/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestMergeSortSorted(t *testing.T) {
	t.Helper()
	taskRunner := NewMergeSortRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/sorted/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}

func TestMergeSortRevers(t *testing.T) {
	t.Helper()
	taskRunner := NewMergeSortRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/revers/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
