package consts

import "fmt"

var (
	UnimplementedError = fmt.Errorf(NotImplementedErrStr)
	EOFError           = fmt.Errorf("%s", "EOF")
)
