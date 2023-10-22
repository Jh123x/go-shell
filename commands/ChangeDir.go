package commands

import (
	"os"

	"github.com/Jh123x/go-shell/consts"
)

type ChangeDirectoryCommand struct {
	BasicCommand
}

func (c ChangeDirectoryCommand) Execute() {
	if len(c.args) != 1 {
		c.PrintErrorString(consts.TooManyArgsErrStr)
		return
	}
	c.PrintError(os.Chdir(c.args[0]))
}

func NewChangeDirectoryCommand(args []string) Command {
	return &ChangeDirectoryCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
