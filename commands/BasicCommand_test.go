package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestNewBasicCommand(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	if basicCommand.inputPipe != os.Stdin {
		t.Errorf("NewBasicCommand() did not set inputPipe to os.Stdin")
	}

	if basicCommand.outputPipe != os.Stdout {
		t.Errorf("NewBasicCommand() did not set outputPipe to os.Stdout")
	}

	if basicCommand.errorPipe != os.Stderr {
		t.Errorf("NewBasicCommand() did not set errorPipe to os.Stderr")
	}
}

func TestExecute(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("TestExecute() could not create pipe")
	}
	basicCommand.errorPipe = w
	basicCommand.Execute()

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("TestExecute() could not read from pipe")
	}

	err_msg := string(buf[:n])
	if err_msg != "error: command not implemented\n" {
		t.Errorf("TestExecute() did not print correct error message, got %s", err_msg)
	}
}

func TestSetPipes(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("TestSetPipes() could not create pipe")
	}

	t.Run("SetInputPipe", func(t *testing.T) {
		basicCommand.SetInputPipe(r)
		if basicCommand.inputPipe != r {
			t.Errorf("TestSetPipes() did not set inputPipe correctly")
		}
	})

	t.Run("SetOutputPipe", func(t *testing.T) {
		basicCommand.SetOutputPipe(w)
		if basicCommand.outputPipe != w {
			t.Errorf("TestSetPipes() did not set outputPipe correctly")
		}
	})

	t.Run("SetErrorPipe", func(t *testing.T) {
		basicCommand.SetErrorPipe(w)
		if basicCommand.errorPipe != w {
			t.Errorf("TestSetPipes() did not set errorPipe correctly")
		}
	})

	// Close all of them
	r.Close()
	w.Close()
}

func TestGetPipes(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("TestGetPipes() could not create pipe")
	}

	basicCommand.inputPipe = r
	basicCommand.outputPipe = w
	basicCommand.errorPipe = w

	t.Run("GetInputPipe", func(t *testing.T) {
		if basicCommand.GetInputPipe() != r {
			t.Errorf("TestGetPipes() did not get inputPipe correctly")
		}
	})

	t.Run("GetOutputPipe", func(t *testing.T) {
		if basicCommand.GetOutputPipe() != w {
			t.Errorf("TestGetPipes() did not get outputPipe correctly")
		}
	})

	t.Run("GetErrorPipe", func(t *testing.T) {
		if basicCommand.GetErrorPipe() != w {
			t.Errorf("TestGetPipes() did not get errorPipe correctly")
		}
	})

	// Close all of them
	r.Close()
	w.Close()
}

// Test Print
func TestPrint(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("TestPrint() could not create pipe")
	}
	basicCommand.outputPipe = w

	basicCommand.Print("test")
	w.Close()
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("TestPrint() could not read from pipe")
	}

	out := string(buf[:n])
	if out != "test\n" {
		t.Errorf("TestPrint() did not print correct message, got '%s'", out)
	}
}

func TestPrintErrorString(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("TestPrintErrorString() could not create pipe")
	}
	basicCommand.errorPipe = w

	basicCommand.PrintErrorString("test")
	w.Close()
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("TestPrintErrorString() could not read from pipe")
	}

	out := string(buf[:n])
	if out != "test\n" {
		t.Errorf("TestPrintErrorString() did not print correct message, got '%s'", out)
	}
}

func TestPrintError(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("TestPrintError() could not create pipe")
	}
	basicCommand.errorPipe = w

	basicCommand.PrintError(fmt.Errorf("test"))
	w.Close()
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("TestPrintError() could not read from pipe")
	}

	out := string(buf[:n])
	if out != "test\n" {
		t.Errorf("TestPrintError() did not print correct message, got '%s'", out)
	}
}

func TestClose(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	// Do not close standard pipes
	basicCommand.Close()
	if basicCommand.inputPipe != os.Stdin {
		t.Errorf("TestClose() closed standard inputPipe")
	}

	if basicCommand.outputPipe != os.Stdout {
		t.Errorf("TestClose() closed standard outputPipe")
	}

	if basicCommand.errorPipe != os.Stderr {
		t.Errorf("TestClose() closed standard errorPipe")
	}
}
