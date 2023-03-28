package commands

import (
	"os"
)

type RemoveFileCommand struct {
	BasicCommand
}

func (r RemoveFileCommand) Execute() {
	if len(r.args) < 1 {
		r.PrintErrorString("rm: missing argument")
		return
	}
	for _, arg := range r.args {
		err := os.Remove(arg)
		if err != nil {
			r.PrintError(err)
		}
	}
	r.Print("Successfully Removed")
}

func NewRemoveFileCommand(args []string) Command {
	return &RemoveFileCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
