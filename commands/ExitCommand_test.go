package commands

import (
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/Jh123x/go-shell/consts"
	"github.com/stretchr/testify/assert"
)

func TestNewExitCommand(t *testing.T) {
	type args struct {
		args []string
		want Command
	}
	tests := []args{
		{
			args: []string{},
			want: &ExitCommand{},
		},
	}
	for _, tt := range tests {
		t.Run("TestNewExitCommand", func(t *testing.T) {
			assert.NotNil(t, NewExitCommand(tt.args))
		})
	}
}

func TestExecuteExitWithArgs(t *testing.T) {

	// Create pipe
	r, w, err := os.Pipe()
	assert.Nil(t, err)

	cmd := NewExitCommand([]string{"test"})
	cmd.SetErrorPipe(w)
	cmd.Execute()
	w.Close()

	// Read from pipe
	res, err := io.ReadAll(r)
	assert.Nil(t, err)
	assert.Equal(t, consts.TooManyArgsErrStr+"\n", string(res))
}

func TestExecuteExit(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		NewExitCommand([]string{}).Execute()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	assert.Nil(t, err)
}
