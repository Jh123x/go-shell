package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	command "commands"
)

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
			out, err2 := command.Execute()
			if err2 != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(out)
		}
	}
}

func parseInput(input string) (string, []string) {
	input = strings.Trim(input, "\n\r\t ")
	args := strings.Split(input, " ")
	return args[0], args[1:]
}

// Execute a command based on a clean input
func inputToCommand(cmd string, args []string) command.Command {
	switch cmd {
	case "dir", "ls":
		return command.NewListDirectoryCommand(args)
	case "cd":
		return command.NewChangeDirectoryCommand(args)
	case "pwd":
		return command.NewCwdCommand(args)
	case "exit":
		return command.NewExitCommand(args)
	case "cat":
		return command.NewPrintFileCommand(args)
	case "rm":
		return command.NewRemoveFileCommand(args)
	default:
		return command.NewDefaultCommand([]string{cmd, args[0]})
	}
}
