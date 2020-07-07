// Copyright (c) 2017 pavelblossom. All rights reserved.

package backend

import (
	"errors"
	"io"
	"os/exec"
)

type Local struct{}

func (b *Local) StartProcess(cmd string, args ...string) (Waiter, io.Writer, io.Reader, io.Reader, error) {
	command := exec.Command(cmd, args...)

	stdin, err := command.StdinPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.New(err.Error() + ". Could not get hold of the PowerShell's stdin stream")
	}

	stdout, err := command.StdoutPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.New(err.Error() + ". Could not get hold of the PowerShell's stdout stream")
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		return nil, nil, nil, nil, errors.New(err.Error() + ". Could not get hold of the PowerShell's stderr stream")
	}

	err = command.Start()
	if err != nil {
		return nil, nil, nil, nil, errors.New(err.Error() + ".Could not spawn PowerShell process")
	}

	return command, stdin, stdout, stderr, nil
}
