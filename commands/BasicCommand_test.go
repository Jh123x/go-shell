package commands

import (
	"fmt"
	"os"
	"testing"

	"github.com/Jh123x/go-shell/consts"
	"github.com/stretchr/testify/assert"
)

func TestNewBasicCommand(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})
	assert.Equal(t, os.Stdin, basicCommand.inputPipe)
	assert.Equal(t, os.Stdout, basicCommand.outputPipe)
	assert.Equal(t, os.Stderr, basicCommand.errorPipe)
}

func TestExecute(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	assert.Nil(t, err)
	basicCommand.errorPipe = w
	basicCommand.Execute()

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	assert.Nil(t, err)

	err_msg := string(buf[:n])
	assert.Equal(t, "error: command not implemented\n", err_msg)
}

func TestSetPipes(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	assert.Nil(t, err)
	defer r.Close()
	defer w.Close()

	t.Run("SetInputPipe", func(t *testing.T) {
		basicCommand.SetInputPipe(r)
		assert.Equal(t, r, basicCommand.inputPipe)
	})

	t.Run("SetOutputPipe", func(t *testing.T) {
		basicCommand.SetOutputPipe(w)
		assert.Equal(t, w, basicCommand.outputPipe)
	})

	t.Run("SetErrorPipe", func(t *testing.T) {
		basicCommand.SetErrorPipe(w)
		assert.Equal(t, w, basicCommand.errorPipe)
	})
}

func TestGetPipes(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	r, w, err := os.Pipe()
	assert.Nil(t, err)
	defer r.Close()
	defer w.Close()

	basicCommand.inputPipe = r
	basicCommand.outputPipe = w
	basicCommand.errorPipe = w

	t.Run("GetInputPipe", func(t *testing.T) {
		assert.Equal(t, r, basicCommand.GetInputPipe())
	})

	t.Run("GetOutputPipe", func(t *testing.T) {
		assert.Equal(t, w, basicCommand.GetOutputPipe())
	})

	t.Run("GetErrorPipe", func(t *testing.T) {
		assert.Equal(t, w, basicCommand.GetErrorPipe())
	})
}

// Test Print
func TestPrint(t *testing.T) {
	tests := map[string]struct {
		input       string
		expectedErr error
		expectedRes string
	}{
		"valid string": {
			input:       "test",
			expectedErr: nil,
			expectedRes: "test\n",
		},
		"empty string": {
			input:       "",
			expectedErr: consts.EOFError,
			expectedRes: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			basicCommand := NewBasicCommand([]string{""})

			r, w, err := os.Pipe()
			assert.Nil(t, err)
			basicCommand.outputPipe = w
			basicCommand.Print(tc.input)
			w.Close()
			buf := make([]byte, 1024)
			n, err := r.Read(buf)
			assert.Equal(t, tc.expectedErr, err)

			out := string(buf[:n])
			assert.Equal(t, tc.expectedRes, out)
		})
	}

}

func TestPrintErrorString(t *testing.T) {

	test := map[string]struct {
		input         string
		expectedWrite string
		expectedErr   error
	}{
		"valid string": {
			input:         "test",
			expectedWrite: "test\n",
			expectedErr:   nil,
		},
		"empty string": {
			input:         "",
			expectedWrite: "",
			expectedErr:   consts.EOFError,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			basicCommand := NewBasicCommand([]string{""})

			r, w, err := os.Pipe()
			assert.Nil(t, err)
			defer r.Close()
			basicCommand.errorPipe = w

			basicCommand.PrintErrorString(tc.input)
			w.Close()
			buf := make([]byte, 1024)
			n, err := r.Read(buf)
			assert.Equal(t, tc.expectedErr, err)

			out := string(buf[:n])
			assert.Equal(t, tc.expectedWrite, out)
		})
	}
}

func TestPrintError(t *testing.T) {
	tests := map[string]struct {
		inputErr    error
		expectedRes string
		expectedErr error
	}{
		"valid error": {
			inputErr:    fmt.Errorf("test"),
			expectedRes: "test\n",
			expectedErr: nil,
		},
		"empty msg": {
			inputErr:    nil,
			expectedRes: "",
			expectedErr: consts.EOFError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			basicCommand := NewBasicCommand([]string{""})

			r, w, err := os.Pipe()
			assert.Nil(t, err)
			basicCommand.errorPipe = w
			basicCommand.PrintError(tc.inputErr)
			w.Close()
			buf := make([]byte, 1024)
			n, err := r.Read(buf)
			assert.Equal(t, tc.expectedErr, err)

			out := string(buf[:n])
			assert.Equal(t, tc.expectedRes, out)
		})
	}

}

func TestClose(t *testing.T) {
	basicCommand := NewBasicCommand([]string{""})

	// Do not close standard pipes
	basicCommand.Close()
	assert.Equal(t, os.Stdin, basicCommand.inputPipe)
	assert.Equal(t, os.Stdout, basicCommand.outputPipe)
	assert.Equal(t, os.Stderr, basicCommand.errorPipe)
}
