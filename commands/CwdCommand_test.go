package commands

import (
	"io"
	"os"
	"testing"

	"github.com/Jh123x/go-shell/consts"
	"github.com/stretchr/testify/assert"
)

// NewCwdCommand returns a new CwdCommand
func TestNewCwdCommand(t *testing.T) {
	// Struct for the tests
	type args struct {
		args []string
		want Command
	}

	// Create the tests
	tests := []args{
		{
			args: []string{},
			want: &CwdCommand{},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(
			"TestNewCwdCommand",
			func(t *testing.T) {
				if got := NewCwdCommand(tt.args); got == nil {
					t.Errorf("NewCwdCommand() returned nil")
				}
			},
		)
	}
}

// Execute executes the CwdCommand
func TestExecuteCwd(t *testing.T) {
	tests := map[string]struct {
		args        []string
		expectedErr string
	}{
		"no args": {
			args: []string{},
		},
		"with args": {
			args:        []string{"test"},
			expectedErr: consts.TooManyArgsErrStr + "\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// Save current dir
			cwd, err := os.Getwd()
			assert.Nil(t, err)

			r, w, err := os.Pipe()
			assert.Nil(t, err)

			//err pipe
			rErr, wErr, err := os.Pipe()
			assert.Nil(t, err)

			cmd := NewCwdCommand(tc.args)
			cmd.SetOutputPipe(w)
			cmd.SetErrorPipe(wErr)
			cmd.Execute()
			cmd.Close()
			assert.Nil(t, err)

			// Check if it matches cwd
			newCwd, err := io.ReadAll(r)
			if len(tc.expectedErr) > 0 {
				rErrBytes, err := io.ReadAll(rErr)
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedErr, string(rErrBytes))
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cwd, string(newCwd)[0:len(newCwd)-1])
			}

		})
	}
}
