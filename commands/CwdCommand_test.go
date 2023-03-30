package commands

import (
	"io/ioutil"
	"os"
	"testing"
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
	// Save current dir
	cwd, err := os.Getwd()

	if err != nil {
		t.Errorf("Could not get current working directory")
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("Could not create pipe")
	}

	cmd := NewCwdCommand([]string{})
	cmd.SetOutputPipe(w)
	cmd.Execute()
	cmd.Close()

	if err != nil {
		t.Errorf("Could not get current working directory")
	}

	// Check if it matches cwd
	newCwd, err := ioutil.ReadAll(r)
	if err != nil {
		t.Errorf("Could not read from pipe")
	}

	cwdCmd := string(newCwd)[0 : len(newCwd)-1]
	if cwdCmd != cwd {
		t.Errorf("Directory did not match, got %s, want %s", cwdCmd, cwd)
	}
}
