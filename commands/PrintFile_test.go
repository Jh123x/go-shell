package commands

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintCommand(t *testing.T) {
	tests := map[string]struct {
		args        []string
		stdin       string
		expectedOut string
		expectedErr string
	}{
		"empty file": {
			args:        []string{"../tests/empty.txt"},
			expectedOut: "",
			expectedErr: "",
		},
		"file with contents": {
			args:        []string{"../tests/one_liner.txt"},
			expectedOut: "this is a line\n",
			expectedErr: "",
		},
		"file not found": {
			args:        []string{"../tests/does_not_exist.txt"},
			expectedOut: "",
			expectedErr: "open ../tests/does_not_exist.txt: The system cannot find the file specified.\n",
		},
		"no args": {
			args:        []string{},
			stdin:       "this is a line\n",
			expectedOut: "this is a line\n",
			expectedErr: "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.Nil(t, err)

			rOut, wOut, err := os.Pipe()
			assert.Nil(t, err)

			rErr, wErr, err := os.Pipe()
			assert.Nil(t, err)

			cmd := NewPrintFileCommand(testCase.args)
			cmd.SetInputPipe(r)
			w.Write([]byte(testCase.stdin))
			w.Close()
			cmd.SetOutputPipe(wOut)
			cmd.SetErrorPipe(wErr)
			cmd.Execute()

			wOut.Close()
			wErr.Close()

			outBytes, err := io.ReadAll(rOut)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedOut, string(outBytes))

			errBytes, err := io.ReadAll(rErr)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedErr, string(errBytes))
		})
	}
}
