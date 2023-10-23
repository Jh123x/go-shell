package commands

import (
	"io"
	"os"
	"testing"

	"github.com/Jh123x/go-shell/consts"
	"github.com/stretchr/testify/assert"
)

func TestNewListDirCommand(t *testing.T) {
	type args struct {
		args []string
		want Command
	}
	tests := []args{
		{
			args: []string{},
			want: &ListDirectoryCommand{},
		},
	}
	for _, tt := range tests {
		t.Run("TestNewListDirCommand", func(t *testing.T) {
			assert.NotNil(t, NewListDirectoryCommand(tt.args))
		})
	}
}

func TestExecuteListDir(t *testing.T) {
	type args struct {
		args           []string
		expectedOut    bool
		expectedErrMsg string
	}
	tests := []args{
		{
			args:        []string{},
			expectedOut: true,
		},
		{
			args:        []string{"."},
			expectedOut: true,
		},
		{
			args:           []string{"../..", ".."},
			expectedOut:    false,
			expectedErrMsg: consts.TooManyArgsErrStr + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(
			"TestExecuteListDir",
			func(t *testing.T) {
				cmd := NewListDirectoryCommand(tt.args)
				assert.NotNil(t, cmd)
				r, w, err := os.Pipe()
				assert.Nil(t, err)
				rErr, wErr, err := os.Pipe()
				assert.Nil(t, err)
				cmd.SetOutputPipe(w)
				cmd.SetErrorPipe(wErr)
				cmd.Execute()
				cmd.Close()

				outputByte, err := io.ReadAll(r)
				assert.Nil(t, err)
				output := string(outputByte)
				assert.Equal(t, tt.expectedOut, len(output) > 0)

				errByte, err := io.ReadAll(rErr)
				assert.Nil(t, err)
				errMsg := string(errByte)
				assert.Equal(t, tt.expectedErrMsg, errMsg)
			},
		)
	}
}
