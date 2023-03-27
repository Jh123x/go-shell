package commands

import (
	"fmt"
	"os"
)

type ChangeDirectoryCommand struct {
	BasicCommand
}

func (c ChangeDirectoryCommand) Execute() (string, error) {
	if len(c.args) < 1 {
		return "", fmt.Errorf("cd: missing argument")
	}
	if len(c.args) > 1 {
		return "", fmt.Errorf("cd: too many arguments")
	}
	return "", os.Chdir(c.args[0])
}

func NewChangeDirectoryCommand(args []string) *ChangeDirectoryCommand {
	return &ChangeDirectoryCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
