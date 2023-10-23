package commands

import (
	"os"
	"strings"

	"github.com/Jh123x/go-shell/consts"
)

type ListDirectoryCommand struct {
	BasicCommand
}

func (l ListDirectoryCommand) Execute() {
	arg := "./"
	if len(l.args) > 1 {
		l.PrintErrorString(consts.TooManyArgsErrStr)
		return
	}
	if len(l.args) == 1 {
		arg = l.args[0]
	}
	entries, err := os.ReadDir(arg)
	if err != nil {
		l.PrintError(err)
		return
	}
	arr := []string{}
	for _, f := range entries {
		arr = append(arr, f.Name())
	}
	l.Print(strings.Join(arr, "\n"))
}

func NewListDirectoryCommand(args []string) Command {
	return &ListDirectoryCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
