package commands

import "os"

type Command interface {
	Execute()
	SetInputPipe(pipe *os.File)
	SetOutputPipe(pipe *os.File)
	SetErrorPipe(pipe *os.File)
}
