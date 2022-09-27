package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		currDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Print(currDir + "> go-shell >> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	args := strings.Split(input, " ")
	switch args[0] {
	case "dir", "ls":
		return ListDirectoryCommand(args)
	case "cd":
		return ChangeDirectoryCommand(args[1])
	case "pwd":
		return GetCurrentDirectoryCommand()
	case "exit":
		os.Exit(0)
		return nil
	default:
		return DefaultCommand(args[0], args[1:])
	}
}

func GetCurrentDirectoryCommand() error {
	currDir, err := os.Getwd()
	fmt.Println(currDir)
	return err
}

func ListDirectoryCommand(args []string) error {
	arg := "./"
	if len(args) > 1 {
		arg = args[1]
	}
	entries, err := os.ReadDir(arg)
	for _, f := range entries {
		fmt.Println(f.Name())
	}
	return err
}

func ChangeDirectoryCommand(dir string) error {
	return os.Chdir(dir)
}

func DefaultCommand(cmd_str string, args []string) error {
	cmd := exec.Command(cmd_str, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
