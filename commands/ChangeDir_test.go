package commands

import (
	"os"
	"strings"
	"testing"

	"github.com/Jh123x/go-shell/consts"
	"github.com/stretchr/testify/assert"
)

func TestNewChangeDirectoryCommand(t *testing.T) {
	// Struct for the tests
	type args struct {
		args []string
		want Command
	}

	// Create the tests
	tests := []args{
		{
			args: []string{".."},
			want: &ChangeDirectoryCommand{},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(
			"TestNewChangeDirectoryCommand",
			func(t *testing.T) {
				assert.NotNil(t, NewChangeDirectoryCommand(tt.args))
			},
		)
	}
}

func TestExecuteChangeDir(t *testing.T) {
	// Save current dir
	cwd, err := os.Getwd()
	assert.Nil(t, err)

	tests := map[string]struct {
		args        []string
		expectedErr string
		expectedCwd string
	}{
		"change dir to itself": {
			args:        []string{"."},
			expectedErr: "",
			expectedCwd: cwd,
		},
		"wrong no of args": {
			args:        []string{"test", "test"},
			expectedErr: consts.TooManyArgsErrStr + "\n",
			expectedCwd: cwd,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() { os.Chdir(cwd) }()
			// Setup pipes
			r, w, err := os.Pipe()
			assert.Nil(t, err)

			// Change to parent dir
			cmd := NewChangeDirectoryCommand(tc.args)
			cmd.SetErrorPipe(w)
			cmd.Execute()
			w.Close()

			// Check if dir changed
			newCwd, err := os.Getwd()
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedCwd, newCwd)

			// Read error stream
			buf := make([]byte, 1024)
			_, err = r.Read(buf)
			if len(tc.expectedErr) == 0 {
				assert.Equal(t, consts.EOFError, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedErr, strings.Trim(string(buf), "\x00"))
			}
		})
	}

}
