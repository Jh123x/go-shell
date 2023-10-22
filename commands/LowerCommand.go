package commands

import (
	"io"
	"strings"

	"github.com/Jh123x/go-shell/consts"
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
				c.PrintErrorString(consts.ReadInputErr)
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
