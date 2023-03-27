package commands

import (
	"fmt"
	"os"
)

type RemoveFileCommand struct {
	BasicCommand
}

func (r RemoveFileCommand) Execute() (string, error) {
	if len(r.args) < 1 {
		return "", fmt.Errorf("rm: missing argument")
	}
	for _, arg := range r.args {
		err := os.Remove(arg)
		if err != nil {
			return "", err
		}
	}
	return "Successfully Removed", nil
}

func NewRemoveFileCommand(args []string) Command {
	return &RemoveFileCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
