package simplesort

import (
	"github.com/alexMolokov/hw-otus-algorithm/system"
	"os"
	"path/filepath"
	"testing"
)

func TestBubbleSortRandom(t *testing.T) {
	t.Helper()
	taskRunner := NewBubbleRunner()
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/random/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
