package commands

import "os"

type Command interface {
	// Execution
	Execute()

	// IO
	SetInputPipe(pipe *os.File)
	GetInputPipe() *os.File
	SetOutputPipe(pipe *os.File)
	GetOutputPipe() *os.File
	SetErrorPipe(pipe *os.File)
	GetErrorPipe() *os.File
}
