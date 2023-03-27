package commands

import (
	"fmt"
	"os/exec"
)

type DefaultCommand struct {
	BasicCommand
}

func (c DefaultCommand) Execute() (string, error) {
	cmd := exec.Command(c.args[0], c.args[1:]...)
	err := cmd.Run()

	if err != nil {
		if errStr := err.Error(); len(errStr) > 19 && errStr[len(errStr)-19:] == "not found in %PATH%" {
			return "", fmt.Errorf("command not found: %s", c.args[0])
		}
	}
	outByte, err := cmd.Output()
	return string(outByte), err
}

func NewDefaultCommand(args []string) *DefaultCommand {
	return &DefaultCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
