package happytickets

import (
	"testing"

	"github.com/alexMolokov/hw-otus-algorithm/system"
)

func TestHappyTickets(t *testing.T) {
	taskRunner := NewTaskRunner(NewTask())

	tester := system.NewTester(t, taskRunner, "./test-files/")
	tester.RunTests()
}
