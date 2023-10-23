package commands

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Jh123x/go-shell/consts"
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
	for _, fileName := range p.args {
		file, err := os.Open(fileName)
		if err != nil {
			p.PrintErrorString(
				fmt.Sprintf(consts.FileNotFoundErrStr, fileName),
			)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			p.Print(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			p.PrintError(scanner.Err())
			return
		}
	}
}

// Constructor
func NewPrintFileCommand(args []string) Command {
	return &PrintFileCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
