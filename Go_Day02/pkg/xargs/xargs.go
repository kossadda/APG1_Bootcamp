// Package xargs provides a utility for building and executing commands
// by appending arguments read from standard input to the given command-line arguments.
package xargs

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/kossadda/APG1_Bootcamp/pkg/message"
)

// XArgs represents the xargs utility, holding the command to be executed and its arguments.
type XArgs struct {
	cmd  string
	args []string
}

// Execute runs the command stored in the XArgs instance with the appended arguments.
// It redirects standard error output to the program's stderr and returns the command's output or an error.
func (x *XArgs) Execute() ([]byte, error) {
	cmd := exec.Command(x.cmd, x.args...)
	cmd.Stderr = os.Stderr

	return cmd.Output()
}

// New initializes a new XArgs instance by reading arguments from standard input
// and appending them to the command and arguments provided as command-line arguments.
// Returns an error if no command is provided or if an error occurs while reading stdin.
func New() (*XArgs, error) {
	var xg XArgs
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		xg.args = append(xg.args, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	secondArgs := os.Args[1:]
	if len(secondArgs) == 0 {
		return nil, message.EmptyCommand()
	}

	xg.cmd = secondArgs[0]
	xg.args = append(secondArgs[1:], xg.args...)

	return &xg, nil
}
