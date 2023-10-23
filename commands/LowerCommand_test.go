package commands

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowerCommand(t *testing.T) {
	tests := map[string]struct {
		stdin       string
		args        []string
		expectedOut string
		expectedErr string
	}{
		"empty stdin": {
			stdin:       "",
			expectedOut: "",
			expectedErr: "",
		},
		"all uppercase": {
			stdin:       "HELLO",
			expectedOut: "hello\n",
			expectedErr: "",
		},
		"all lowercase": {
			stdin:       "hello",
			expectedOut: "hello\n",
			expectedErr: "",
		},
		"mixed case": {
			stdin:       "HeLlO",
			expectedOut: "hello\n",
			expectedErr: "",
		},
		"args": {
			stdin:       "",
			args:        []string{"HelLo", "woRld"},
			expectedOut: "hello world\n",
			expectedErr: "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			//stdin
			r, w, err := os.Pipe()
			assert.Nil(t, err)

			//stdout
			rOut, wOut, err := os.Pipe()
			assert.Nil(t, err)

			//stderr
			rErr, wErr, err := os.Pipe()
			assert.Nil(t, err)

			cmd := NewLowerCommand(testCase.args)
			cmd.SetInputPipe(r)
			cmd.SetOutputPipe(wOut)
			cmd.SetErrorPipe(wErr)

			_, err = w.Write([]byte(testCase.stdin))
			assert.Nil(t, err)
			w.Close()

			cmd.Execute()
			wOut.Close()
			wErr.Close()

			stdout, err := io.ReadAll(rOut)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedOut, string(stdout))

			stderr, err := io.ReadAll(rErr)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedErr, string(stderr))
		})
	}
}
