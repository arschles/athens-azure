package run

import (
	"bytes"
	"os/exec"
	"strings"
)

// Command runs the given command c with args args, and returns STDOUT or an error
func Command(c string, args ...string) (string, error) {
	cmd := exec.Command(c, args...)
	var stdout bytes.Buffer
	// var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return "", err
	}
	output := string(stdout.Bytes())
	ret := strings.TrimRight(output, "\n")
	return ret, nil
}
