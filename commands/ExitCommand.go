package commands

import (
	"os"

	"github.com/Jh123x/go-shell/consts"
)

type ExitCommand struct {
	BasicCommand
}

func (c ExitCommand) Execute() {
	if len(c.args) > 0 {
		c.PrintErrorString(consts.TooManyArgsErrStr)
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
