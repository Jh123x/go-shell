package commands

import "os"

type BasicCommand struct {
	inputPipe  *os.File
	outputPipe *os.File
	errorPipe  *os.File
	args       []string
}

func (c *BasicCommand) SetInputPipe(pipe *os.File) {
	c.inputPipe = pipe
}

func (c *BasicCommand) SetOutputPipe(pipe *os.File) {
	c.outputPipe = pipe
}

func (c *BasicCommand) SetErrorPipe(pipe *os.File) {
	c.errorPipe = pipe
}

func (c *BasicCommand) Execute() (string, error) {
	return "", nil
}

func NewBasicCommand(args []string) BasicCommand {
	return BasicCommand{
		inputPipe:  os.Stdin,
		outputPipe: os.Stdout,
		errorPipe:  os.Stderr,
		args:       args,
	}
}
