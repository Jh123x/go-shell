package main

import (
	"bufio"
	"commands"
	"fmt"
	"os"
	"strings"
)

var cmdMap = map[string]func(args []string) commands.Command{
	"dir":  commands.NewListDirectoryCommand,
	"ls":   commands.NewListDirectoryCommand,
	"cd":   commands.NewChangeDirectoryCommand,
	"pwd":  commands.NewCwdCommand,
	"exit": commands.NewExitCommand,
	"cat":  commands.NewPrintFileCommand,
	"rm":   commands.NewRemoveFileCommand,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		currDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Print(currDir + "$$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		cmds := strings.Split(input, "&&")
		for _, cmd := range cmds {
			cleanedCmd, args := parseInput(cmd)
			command := inputToCommand(cleanedCmd, args)
			command.Execute()
		}
	}
}

func parseInput(input string) (string, []string) {
	input = strings.Trim(input, "\n\r\t ")
	args := strings.Split(input, " ")
	return args[0], args[1:]
}

// Execute a command based on a clean input
func inputToCommand(cmd string, args []string) commands.Command {
	cmdFunc, isFound := cmdMap[cmd]
	if !isFound {
		return commands.NewDefaultCommand(cmd, args)
	}
	return cmdFunc(args)
}
