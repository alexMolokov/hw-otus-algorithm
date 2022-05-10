package system

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type ITaskRunner interface {
	Run(in []string) string
}

type tester struct {
	TaskRunner ITaskRunner
	Path       string
	T          *testing.T
}

func (ts *tester) RunTests() {
	tNum := 0
	sep := "\n"

	if runtime.GOOS == "windows" {
		sep = "\r\n"
	}

	for {

		inFile := fmt.Sprintf(ts.Path+"test.%d.in", tNum)
		outFile := fmt.Sprintf(ts.Path+"test.%d.out", tNum)
		if _, err := os.Stat(inFile); errors.Is(err, os.ErrNotExist) {
			break
		}
		if _, err := os.Stat(outFile); errors.Is(err, os.ErrNotExist) {
			break
		}

		ts.T.Run(fmt.Sprintf("test num %d", tNum), func(t *testing.T) {
			dataIn, err := ioutil.ReadFile(inFile)
			require.NoError(t, err, fmt.Sprintf("Can't read file %s", inFile))
			dataOut, err := ioutil.ReadFile(outFile)
			require.NoError(t, err, fmt.Sprintf("Can't read file %s", outFile))

			data := strings.Split(string(dataIn), sep)
			for k, v := range data {
				data[k] = strings.Trim(v, " "+sep)
			}

			resultTest := ts.TaskRunner.Run(data)
			resultExpected := strings.Trim(string(dataOut), " "+sep)
			require.Equal(t, resultExpected, resultTest, fmt.Sprintf("Must equal inFile=%s, outFile=%s", inFile, outFile))
			fmt.Printf(" in %#v, out = %s, excepted = %s", data, resultTest, resultExpected)
		})

		tNum++
	}

	require.Greater(ts.T, tNum, 0, "Must be more then zero tests")
}

func NewTester(t *testing.T, task ITaskRunner, path string) *tester {
	return &tester{
		TaskRunner: task,
		Path:       path,
		T:          t,
	}
}
