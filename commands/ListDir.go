package commands

import (
	"os"
	"strings"
)

type ListDirectoryCommand struct {
	BasicCommand
}

func (l ListDirectoryCommand) Execute() (string, error) {
	arg := "./"
	if len(l.args) > 1 {
		arg = l.args[0]
	}
	entries, err := os.ReadDir(arg)
	arr := []string{}
	for _, f := range entries {
		arr = append(arr, f.Name())
	}
	return strings.Join(arr, "\n"), err
}

func NewListDirectoryCommand(args []string) Command {
	return &ListDirectoryCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
