package commands

import (
	"fmt"
	"os"
)

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

func (c *BasicCommand) Execute() {
	c.PrintError(fmt.Errorf("error: command not implemented"))
}

func (c *BasicCommand) Print(message string) {
	if len(message) == 0 {
		return
	}
	lenWritten, err := fmt.Fprintln(c.outputPipe, message)
	if err != nil {
		c.PrintError(err)
		return
	}

	if lenWritten < len(message) {
		c.PrintErrorString(fmt.Sprintf("error: could not write all bytes to output pipe (%d written, %d message)", lenWritten, len(message)))
	}
}

func (c *BasicCommand) PrintErrorString(errorMessage string) {
	if len(errorMessage) == 0 {
		return
	}
	c.PrintError(fmt.Errorf(errorMessage))
}

func (c *BasicCommand) PrintError(errorMessage error) {
	if errorMessage == nil {
		return
	}
	fmt.Fprintln(c.errorPipe, errorMessage)
}

func (c *BasicCommand) GetInputPipe() *os.File {
	return c.inputPipe
}

func (c *BasicCommand) GetOutputPipe() *os.File {
	return c.outputPipe
}

func (c *BasicCommand) GetErrorPipe() *os.File {
	return c.errorPipe
}

func NewBasicCommand(args []string) BasicCommand {
	return BasicCommand{
		inputPipe:  os.Stdin,
		outputPipe: os.Stdout,
		errorPipe:  os.Stderr,
		args:       args,
	}
}
