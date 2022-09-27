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
		fmt.Print(currDir + "$$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		cmd, args := parseInput(input)
		out, err2 := execInput(cmd, args)
		if err2 != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(out)
	}
}

func parseInput(input string) (string, []string) {
	input = strings.Trim(input, "\n\r\t ")
	args := strings.Split(input, " ")
	return args[0], args[1:]
}

// Execute a command based on a clean input
func execInput(cmd string, args []string) (string, error) {
	switch cmd {
	case "dir", "ls":
		return listDirectoryCommand(args)
	case "cd":
		return changeDirectoryCommand(args)
	case "pwd":
		return os.Getwd()
	case "exit":
		os.Exit(0)
		return "", nil
	case "cat":
		return readFileCommand(args)
	case "rm":
		return removeFileCommand(args)
	default:
		return defaultCommand(cmd, args)
	}
}

func changeDirectoryCommand(args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("cd: missing argument")
	}
	return "", os.Chdir(args[0])
}

func removeFileCommand(args []string) (string, error) {
	for _, arg := range args {
		err := os.Remove(arg)
		if err != nil {
			return "", err
		}
	}
	return "Successfully Removed", nil
}

func readFileCommand(args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("cat: missing argument")
	}
	file, err := os.Open(args[0])
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

func listDirectoryCommand(args []string) (string, error) {
	arg := "./"
	if len(args) > 1 {
		arg = args[0]
	}
	entries, err := os.ReadDir(arg)
	arr := []string{}
	for _, f := range entries {
		arr = append(arr, f.Name())
	}
	return strings.Join(arr, "\n"), err
}

func defaultCommand(cmd_str string, args []string) (string, error) {
	cmd := exec.Command(cmd_str, args...)
	err := cmd.Run()

	if err != nil {
		if errStr := err.Error(); len(errStr) > 19 && errStr[len(errStr)-19:] == "not found in %PATH%" {
			return "", fmt.Errorf("command not found: %s", cmd_str)
		}
	}
	outByte, err := cmd.Output()
	return string(outByte), err
}
