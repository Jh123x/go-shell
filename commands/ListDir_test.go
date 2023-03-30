package commands

import (
	"io/ioutil"
	"os"
	"testing"
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
		t.Run(
			"TestNewListDirCommand",
			func(t *testing.T) {
				if got := NewListDirectoryCommand(tt.args); got == nil {
					t.Errorf("NewListDirCommand() returned nil")
				}
			},
		)
	}
}

func TestExecuteListDir(t *testing.T) {
	type args struct {
		args []string
		want string
	}
	tests := []args{
		{
			args: []string{},
		},
		{
			args: []string{".."},
		},
	}
	for _, tt := range tests {
		t.Run(
			"TestExecuteListDir",
			func(t *testing.T) {
				cmd := NewListDirectoryCommand(tt.args)
				if cmd == nil {
					t.Errorf("NewListDirCommand() returned nil")
				}
				r, w, err := os.Pipe()
				if err != nil {
					t.Errorf("Could not create pipe")
				}
				cmd.SetOutputPipe(w)
				cmd.Execute()
				cmd.Close()

				outputByte, err := ioutil.ReadAll(r)
				if err != nil {
					t.Errorf("Could not read from pipe")
				}
				output := string(outputByte)
				if output == "" {
					t.Errorf("Output was empty")
				}
			},
		)
	}
}
