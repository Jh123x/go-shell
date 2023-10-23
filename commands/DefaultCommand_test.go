package commands

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Jh123x/go-shell/consts"
	"github.com/stretchr/testify/assert"
)

func TestDefaultCommandExecute(t *testing.T) {
	tests := map[string]struct {
		cmd         string
		args        []string
		expectedOut string
		expectedErr string
	}{
		"missing command": {
			cmd:         "",
			args:        []string{},
			expectedOut: "",
			expectedErr: consts.MissingArgsErrStr + "\n",
		},
		"command not found": {
			cmd:         "notfound",
			args:        []string{},
			expectedOut: "",
			expectedErr: fmt.Sprintf(consts.CmdNotFoundErrStr+"\n", "notfound"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// err Pipe
			rErr, wErr, err := os.Pipe()
			assert.Nil(t, err)
			// out Pipe
			r, w, err := os.Pipe()
			assert.Nil(t, err)

			cmd := NewDefaultCommand(test.cmd, test.args)
			cmd.SetErrorPipe(wErr)
			cmd.SetOutputPipe(w)
			cmd.Execute()

			w.Close()
			wErr.Close()

			stdout, err := io.ReadAll(r)
			assert.Nil(t, err)
			assert.Equal(t, test.expectedOut, string(stdout))

			stderr, err := io.ReadAll(rErr)
			assert.Nil(t, err)
			assert.Equal(t, test.expectedErr, string(stderr))
		})
	}
}
