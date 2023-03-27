package commands

import "os"

type Command interface {
	Execute() (string, error)
	SetInputPipe(pipe *os.File)
	SetOutputPipe(pipe *os.File)
	SetErrorPipe(pipe *os.File)
}
