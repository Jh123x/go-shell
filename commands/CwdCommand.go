package commands

import (
	"os"
)

type CwdCommand struct {
	BasicCommand
}

func (c CwdCommand) Execute() {
	if len(c.args) > 0 {
		c.PrintErrorString("cwd: too many arguments")
	}
	cwd, err := os.Getwd()
	c.PrintError(err)
	c.Print(cwd)
}

func NewCwdCommand(args []string) Command {
	return &CwdCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
