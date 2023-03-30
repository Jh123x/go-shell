package commands

import (
	"io"
	"strings"
)

type LowerCommand struct {
	BasicCommand
}

func (c LowerCommand) Execute() {

	// If there are no args, read from stdin
	if len(c.args) < 1 {

		buffer := make([]byte, 1024)
		for {
			n, err := c.GetInputPipe().Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				c.PrintErrorString("Error reading from input")
				c.PrintError(err)
				break
			}
			c.Print(strings.ToLower(string(buffer[:n])))
			// If reach EOF break
			buffer = make([]byte, 1024)

		}

		return
	}
	// Join args with space
	c.Print(strings.ToUpper(strings.Join(c.args, " ")))

}

func NewLowerCommand(args []string) Command {
	return &LowerCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
