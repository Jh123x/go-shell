package commands

import (
	"testing"
)

// TestNewCommandManager tests the NewCommandManager function
func TestNewCommandManager(t *testing.T) {
	cm := NewCommandManager()
	if cm == nil {
		t.Errorf("NewCommandManager() returned nil")
	}
}

// TestGetCommand tests the GetCommand function
func TestGetCommand(t *testing.T) {
	cm := NewCommandManager()
	for cmd, _ := range mappedCommands {
		_, ok := cm.GetCommand(cmd)
		if !ok {
			t.Errorf("GetCommand(%s) returned false", cmd)
		}
	}
}
