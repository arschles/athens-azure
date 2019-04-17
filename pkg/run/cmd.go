package run

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type Result struct {
	Out string
	Err string
}

// Command runs the given command c with args args, and returns STDOUT or an error
func Command(c string, args ...string) (*Result, error) {
	cmd := exec.Command(c, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, errors.WithMessage(err, string(stderr.Bytes()))
	}
	ret := &Result{
		Out: strings.TrimRight(string(stdout.Bytes()), "\n"),
		Err: strings.TrimRight(string(stderr.Bytes()), "\n"),
	}
	return ret, nil
}
