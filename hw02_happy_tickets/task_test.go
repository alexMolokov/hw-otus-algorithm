package happytickets

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestHappyTickets(t *testing.T) {
	t.Helper()
	taskRunner := NewTaskRunner(NewTask())
	pwd, _ := os.Getwd()
	p := pwd + filepath.FromSlash("/test-files/")
	tester := system.NewTester(t, taskRunner, p)
	tester.RunTests()
}
