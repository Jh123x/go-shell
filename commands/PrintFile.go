package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PrintFileCommand struct {
	BasicCommand
}

func (p PrintFileCommand) Execute() (string, error) {
	if len(p.args) < 1 {
		return "", fmt.Errorf("cat: missing argument")
	}
	file, err := os.Open(p.args[0])
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := []string{}
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	return strings.Join(arr, "\n"), scanner.Err()
}

// Constructor
func NewPrintFileCommand(args []string) *PrintFileCommand {
	return &PrintFileCommand{
		BasicCommand: NewBasicCommand(args),
	}
}
