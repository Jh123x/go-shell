package commands

import (
	"reflect"
	"testing"
)

// Create tests
func TestNewCommandManager(t *testing.T) {
	cm := NewCommandManager()
	if cm == nil {
		t.Errorf("NewCommandManager() returned nil")
	}
}

func TestGetCommand(t *testing.T) {

	// Struct for the tests
	type args struct {
		name string
		arg  string
		want Command
	}

	emptyArgs := []string{""}
	tests := make([]args, 0, len(mappedCommands))
	cm := NewCommandManager()

	// Create the tests
	for cmd, _ := range mappedCommands {
		tests = append(tests, args{
			name: cmd,
			arg:  cmd,
			want: mappedCommands[cmd](emptyArgs),
		})
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				if cmd, ok := cm.GetCommand(tt.name); cmd == nil || !ok {
					t.Errorf("GetCommand() returned nil or false")
				} else {
					result := cmd(emptyArgs)
					// Check if commands are the same type
					if reflect.TypeOf(result) != reflect.TypeOf(tt.want) {
						t.Errorf("GetCommand() returned wrong command type")
					}
				}
			},
		)
	}
}
