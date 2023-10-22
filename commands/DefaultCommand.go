package commands

import (
	"fmt"
	"os/exec"

	"github.com/Jh123x/go-shell/consts"
)

type DefaultCommand struct {
	BasicCommand
}

func (c DefaultCommand) Execute() {
	cmd := exec.Command(c.args[0], c.args[1:]...)
	cmd.Stdout = c.outputPipe
	cmd.Stderr = c.errorPipe
	cmd.Stdin = c.inputPipe

	err := cmd.Run()

	if err != nil {
		if errStr := err.Error(); len(errStr) > 19 && errStr[len(errStr)-19:] == consts.NotFoundInPathPartialString {
			c.PrintErrorString(fmt.Sprintf(consts.CmdNotFoundErr, c.args[0]))
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
