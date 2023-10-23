package commands

import (
	"bufio"
	"io"
	"os"
)

type PrintFileCommand struct {
	BasicCommand
}

func (p PrintFileCommand) Execute() {

	// Echo input to output
	if len(p.args) < 1 {
		io.Copy(p.outputPipe, p.inputPipe)
		return
	}
	file, err := os.Open(p.args[0])
	if err != nil {
		p.PrintError(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.Print(scanner.Text())
	}
	p.PrintError(scanner.Err())
}

// Constructor
func NewPrintFileCommand(args []string) Command {
	return &PrintFileCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
