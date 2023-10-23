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

	if len(c.args) > 0 {
		// Join args with space
		c.Print(strings.ToLower(strings.Join(c.args, " ")))
		return
	}
	// If there are no args, read from stdin
	buffer := make([]byte, 1024)
	for {
		n, err := c.GetInputPipe().Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			c.PrintErrorString(consts.ReadInputErrStr)
			c.PrintError(err)
			break
		}
		c.Print(strings.ToLower(string(buffer[:n])))
		// If reach EOF break
		buffer = make([]byte, 1024)
	}
}

func NewLowerCommand(args []string) Command {
	return &LowerCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
