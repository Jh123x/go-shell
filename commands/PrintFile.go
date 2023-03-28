package commands

import (
	"bufio"
	"os"
	"strings"
)

type PrintFileCommand struct {
	BasicCommand
}

func (p PrintFileCommand) Execute() {
	if len(p.args) < 1 {
		p.PrintErrorString("cat: missing argument")
	}
	file, err := os.Open(p.args[0])
	if err != nil {
		p.PrintError(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := []string{}
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	p.Print(strings.Join(arr, "\n"))
	err = scanner.Err()
	p.PrintError(err)
}

// Constructor
func NewPrintFileCommand(args []string) Command {
	return &PrintFileCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
