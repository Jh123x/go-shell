package commands

import (
	"os"
	"testing"
)

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
