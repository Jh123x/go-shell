package commands

import (
	"os"

	"github.com/Jh123x/go-shell/consts"
)

type CwdCommand struct {
	BasicCommand
}

func (c CwdCommand) Execute() {
	if len(c.args) > 0 {
		c.PrintErrorString(consts.TooManyArgsErr)
		return
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
