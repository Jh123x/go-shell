package commands

import (
	"os"
	"testing"
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
				if got := NewChangeDirectoryCommand(tt.args); got == nil {
					t.Errorf("NewChangeDirectoryCommand() returned nil")
				}
			},
		)
	}
}

func TestExecuteChangeDir(t *testing.T) {
	// Save current dir
	cwd, err := os.Getwd()

	if err != nil {
		t.Errorf("Could not get current working directory")
	}

	// Change to parent dir
	cmd := NewChangeDirectoryCommand([]string{".."})
	cmd.Execute()

	// Check if dir changed
	newCwd, err := os.Getwd()

	if err != nil {
		t.Errorf("Could not get current working directory")
	}

	if newCwd == cwd {
		t.Errorf("Directory did not change")
	}

}
