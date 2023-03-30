package commands

import (
	"os"
	"os/exec"
	"testing"
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
		t.Run(
			"TestNewExitCommand",
			func(t *testing.T) {
				if got := NewExitCommand(tt.args); got == nil {
					t.Errorf("NewExitCommand() returned nil")
				}
			},
		)
	}
}

func TestExecuteExit(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		NewExitCommand([]string{}).Execute()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if err != nil {
		t.Errorf("ExitCommand did not exit correctly.")
	}
}
