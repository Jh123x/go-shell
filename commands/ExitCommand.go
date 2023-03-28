package commands

import (
	"os"
)

type ExitCommand struct {
	BasicCommand
}

func (c ExitCommand) Execute() {
	if len(c.args) > 0 {
		c.PrintErrorString("exit: too many arguments")
		return
	}
	// Function ends here
	os.Exit(0)
}

func NewExitCommand(args []string) Command {
	return &ExitCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
