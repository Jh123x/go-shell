package main

import (
	"commands"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	// Struct for the tests
	type args struct {
		input     string
		want_cmd  string
		want_args []string
	}

	tests := []args{
		{
			input:     "echo hello world",
			want_cmd:  "echo",
			want_args: []string{"hello", "world"},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(
			tt.input,
			func(t *testing.T) {
				if cmd, args := parseInput(tt.input); cmd != tt.want_cmd || !reflect.DeepEqual(args, tt.want_args) {
					t.Errorf("parseInput() returned wrong command or arguments")
				}
			},
		)
	}
}

func TestInputToCommand(t *testing.T) {
	// Struct for the tests
	type args struct {
		name      string
		cmd       string
		args      []string
		want_type reflect.Type
	}

	tests := []args{
		{
			name:      "DefaultCommand",
			cmd:       "echo",
			args:      []string{"hello", "world"},
			want_type: reflect.TypeOf(commands.NewDefaultCommand("a", []string{"hello", "world"})),
		},
	}

	// Setup test based on CommandManager
	cmdManager := commands.NewCommandManager()
	for name, cmd := range cmdManager.GetCommandMap() {
		tests = append(tests, args{
			name:      name,
			cmd:       name,
			args:      []string{"hello", "world"},
			want_type: reflect.TypeOf(cmd([]string{})),
		})
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				// Type of command
				if cmd := inputToCommand(tt.cmd, tt.args); cmd == nil || reflect.TypeOf(cmd) != tt.want_type {
					t.Errorf("inputToCommand() returned wrong type. Got %v, want %v", reflect.TypeOf(cmd), tt.want_type)
				}
			},
		)
	}
}
