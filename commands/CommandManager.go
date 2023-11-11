package commands

type CommandManager struct {
	mappedCommands map[string]func(args []string) Command
}

// GetCommand returns a command from the CommandManager
func (c *CommandManager) GetCommand(name string) (func(args []string) Command, bool) {
	cmd, ok := c.mappedCommands[name]
	return cmd, ok
}

func (c *CommandManager) GetCommandMap() map[string]func(args []string) Command {
	return c.mappedCommands
}

var mappedCommands = map[string]func(args []string) Command{
	"dir":   NewListDirectoryCommand,
	"ls":    NewListDirectoryCommand,
	"cd":    NewChangeDirectoryCommand,
	"pwd":   NewCwdCommand,
	"exit":  NewExitCommand,
	"cat":   NewPrintFileCommand,
	"rm":    NewRemoveFileCommand,
	"up":    NewUpperCommand,
	"dn":    NewLowerCommand,
	"clear": NewClearCommand,
}

func NewCommandManager() *CommandManager {
	return &CommandManager{
		mappedCommands: mappedCommands,
	}
}
