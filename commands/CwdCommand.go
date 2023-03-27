package commands

import (
	"fmt"
	"os"
)

type CwdCommand struct {
	BasicCommand
}

func (c CwdCommand) Execute() (string, error) {
	if len(c.args) > 0 {
		return "", fmt.Errorf("cwd: too many arguments")
	}
	return os.Getwd()
}

func NewCwdCommand(args []string) *CwdCommand {
	return &CwdCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
