package commands

import (
	"fmt"
	"os"
)

type ExitCommand struct {
	BasicCommand
}

func (c ExitCommand) Execute() (string, error) {
	if len(c.args) > 0 {
		return "", fmt.Errorf("exit: too many arguments")
	}
	// Function ends here
	os.Exit(0)
	return "", nil
}

func NewExitCommand(args []string) *ExitCommand {
	return &ExitCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
