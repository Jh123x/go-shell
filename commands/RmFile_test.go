package commands

import (
	"io"
	"os"
	"testing"

	"github.com/Jh123x/go-shell/consts"
	"github.com/stretchr/testify/assert"
)

func TestRmFile(t *testing.T) {
	tests := map[string]struct {
		args            []string
		createdFileName string
		expectedOut     string
		expectedErr     string
	}{
		"file not found": {
			args:            []string{"../tests/does_not_exist.txt"},
			createdFileName: "",
			expectedOut:     "",
			expectedErr:     "remove ../tests/does_not_exist.txt: The system cannot find the file specified.\n",
		},
		"file found": {
			args:            []string{"../tests/test_file.txt"},
			createdFileName: "../tests/test_file.txt",
			expectedOut:     "Successfully Removed\n",
			expectedErr:     "",
		},
		"no args": {
			args:            []string{},
			createdFileName: "",
			expectedOut:     "",
			expectedErr:     consts.MissingArgsErrStr + "\n",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			createTmpFile(t, testCase.createdFileName)
			r, w, err := os.Pipe()
			assert.Nil(t, err)
			rErr, wErr, err := os.Pipe()
			assert.Nil(t, err)

			cmd := NewRemoveFileCommand(testCase.args)
			cmd.SetOutputPipe(w)
			cmd.SetErrorPipe(wErr)
			cmd.Execute()

			w.Close()
			wErr.Close()

			res, err := io.ReadAll(r)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedOut, string(res))

			res, err = io.ReadAll(rErr)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedErr, string(res))
		})
	}
}

func createTmpFile(t *testing.T, fileName string) {
	if len(fileName) == 0 {
		return
	}
	f, err := os.Create(fileName)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
}
