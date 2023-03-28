package commands

import (
	"os"
)

type ChangeDirectoryCommand struct {
	BasicCommand
}

func (c ChangeDirectoryCommand) Execute() {
	if len(c.args) < 1 {
		c.PrintErrorString("cd: too few arguments")
		return
	}
	if len(c.args) > 1 {
		c.PrintErrorString("cd: too many arguments")
		return
	}
	c.PrintError(os.Chdir(c.args[0]))
}

func NewChangeDirectoryCommand(args []string) Command {
	return &ChangeDirectoryCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
