package commands

import (
	"fmt"
	"os"

	"github.com/Jh123x/go-shell/consts"
)

type RemoveFileCommand struct {
	BasicCommand
}

func (r RemoveFileCommand) Execute() {
	if len(r.args) < 1 {
		r.PrintErrorString(consts.MissingArgsErrStr)
		return
	}
	for _, arg := range r.args {
		err := os.Remove(arg)
		if err != nil {
			r.PrintErrorString(
				fmt.Sprintf(consts.FileNotFoundErrStr, arg),
			)
			return
		}
	}
	r.Print(consts.FileRemoveSuccess)
}

func NewRemoveFileCommand(args []string) Command {
	return &RemoveFileCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
