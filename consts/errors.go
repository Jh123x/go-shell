package consts

import "fmt"

var (
	UnimplementedError = fmt.Errorf(NotImplementedErr)
	EOFError           = fmt.Errorf("%s", "EOF")
)
