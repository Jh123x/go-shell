package commands

import (
	"github.com/Jh123x/go-shell/consts"
)

type ClearCommand struct {
	BasicCommand
}

func (c ClearCommand) Execute() {
	if len(c.args) != 0 {
		c.PrintErrorString(consts.TooManyArgsErrStr)
		return
	}
	c.Print("\x1b\x5b\x48\x1b\x5b\x32\x4a")
}

func NewClearCommand(args []string) Command {
	return &ClearCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
