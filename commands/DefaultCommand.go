package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Jh123x/go-shell/consts"
)

type DefaultCommand struct {
	BasicCommand
}

func (c DefaultCommand) Execute() {
	if len(c.args) == 0 || c.args[0] == "" {
		c.PrintErrorString(consts.MissingArgsErrStr)
		return
	}
	cmd := exec.Command(c.args[0], c.args[1:]...)
	cmd.Stdout = c.outputPipe
	cmd.Stderr = c.errorPipe
	cmd.Stdin = c.inputPipe

	err := cmd.Run()

	if err != nil {
		if errStr := err.Error(); len(errStr) > 19 && strings.Contains(errStr, consts.NotFoundInPathPartialStr) {
			c.PrintErrorString(fmt.Sprintf(consts.CmdNotFoundErrStr, c.args[0]))
			return
		}
		c.PrintError(err)
		return
	}
}

func NewDefaultCommand(cmd string, args []string) Command {
	args = append([]string{cmd}, args...)
	return &DefaultCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
