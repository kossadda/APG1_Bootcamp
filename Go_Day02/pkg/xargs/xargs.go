package xargs

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
)

type XArgs struct {
	cmd  string
	args []string
}

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
		return nil, errors.New("input command name for xargs")
	}

	xg.cmd = secondArgs[0]
	xg.args = append(secondArgs[1:], xg.args...)

	return &xg, nil
}

func (x *XArgs) Execute() ([]byte, error) {
	cmd := exec.Command(x.cmd, x.args...)
	cmd.Stderr = os.Stderr

	return cmd.Output()
}
