package commands

import (
	"io"
	"os"
	"testing"

	"github.com/Jh123x/go-shell/consts"
	"github.com/stretchr/testify/assert"
)

func TestNewClearCommand(t *testing.T) {
	assert.NotNil(t, NewClearCommand([]string{}))
}

func TestExecuteClearCmd(t *testing.T) {
	tests := map[string]struct {
		expectedErr string
		isSuccess   bool
	}{
		"clear with args error": {
			expectedErr: consts.TooManyArgsErrStr + "\n",
			isSuccess:   false,
		},
		"clear with no args": {
			expectedErr: "",
			isSuccess:   true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewClearCommand([]string{})

			// Setup pipes
			r, w, err := os.Pipe()
			assert.Nil(t, err)
			c.SetOutputPipe(w)

			// Setup err pipes
			rErr, wErr, err := os.Pipe()
			assert.Nil(t, err)
			c.SetErrorPipe(wErr)

			c.Execute()
			errMsg, err := io.ReadAll(rErr)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedErr, string(errMsg))
			if !tc.isSuccess {
				return
			}
			msg, err := io.ReadAll(r)
			assert.Nil(t, err)
			assert.Equal(t, "\x1b\x5b\x48\x1b\x5b\x32\x4a", string(msg))
		})
	}
}
