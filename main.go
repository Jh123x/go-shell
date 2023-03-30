package main

import (
	"bufio"
	"commands"
	"fmt"
	"os"
	"strings"
	"sync"
)

var cmdManager = commands.NewCommandManager()

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

		// Not running in parallel because order matters here.
		cmds := strings.Split(input, "&&")
		for _, cmd := range cmds {
			processSubCmd(cmd)
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
	cmdFunc, isFound := cmdManager.GetCommand(cmd)
	if !isFound {
		return commands.NewDefaultCommand(cmd, args)
	}
	return cmdFunc(args)
}

func processSubCmd(cmd string) {
	subCmds := strings.Split(cmd, "|")

	// Make cmd array
	cmdArray := make([]commands.Command, 0, len(subCmds))
	for _, pipedCmds := range subCmds {
		cleanedCmd, args := parseInput(pipedCmds)
		command := inputToCommand(cleanedCmd, args)
		cmdArray = append(cmdArray, command)
	}

	// Setup pipes
	currOutput := os.Stdin
	for _, command := range cmdArray[:len(cmdArray)-1] {
		r, w, err := os.Pipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		// Pipe current output to w
		command.SetInputPipe(currOutput)
		command.SetOutputPipe(w)
		currOutput = r
	}

	// Set final pipe to os.Stdout
	lastCmd := cmdArray[len(cmdArray)-1]
	lastCmd.SetInputPipe(currOutput)

	// Execute commands in parallel and waits for last one
	mux := sync.Mutex{}
	mux.Lock()
	for idx, command := range cmdArray {
		go func(command commands.Command, idx int, mux *sync.Mutex) {
			command.Execute()
			command.Close()
			if idx == len(cmdArray)-1 {
				mux.Unlock()
			}
		}(command, idx, &mux)
	}

	// Wait for last command to finish
	mux.Lock()
}
