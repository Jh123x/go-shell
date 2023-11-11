package commands

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Create tests
func TestNewCommandManager(t *testing.T) {
	defer func() { assert.Nil(t, recover()) }()
	cm := NewCommandManager()
	assert.NotNil(t, cm)
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
	for cmd := range mappedCommands {
		tests = append(tests, args{
			name: cmd,
			arg:  cmd,
			want: mappedCommands[cmd](emptyArgs),
		})
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, ok := cm.GetCommand(tt.name)
			assert.NotNil(t, cmd)
			assert.True(t, ok)
			result := cmd(emptyArgs)
			assert.Equal(t, reflect.TypeOf(result), reflect.TypeOf(tt.want))
		})
	}
}

func TestGetCommandMap(t *testing.T) {
	cm := NewCommandManager()
	assert.NotNil(t, cm)
	cmdMap := cm.GetCommandMap()

	// Check if map is equal to mappedCommands
	assert.Equal(t, cmdMap, mappedCommands)
}
