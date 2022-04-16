package osutils

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// CreateDir creates dirs even if the parent dir doesn't exist. (It uses 750 permissions)
func CreateDir(path string) error {
	err := os.MkdirAll(path, 0750)
	if err != nil {
		return fmt.Errorf("couldn't create directory: %s. Error: %w", path, err)
	}
	return nil
}

// NewProcess starts a new process and waits for it to exit.
//
// It accepts multiple arguments.
func NewProcess(groupProcess bool, cmd string, args ...string) (*exec.Cmd, error) {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	// https://stackoverflow.com/a/29552044
	if groupProcess {
		c.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}
	}

	errStart := c.Start()
	if errStart != nil {
		return nil, fmt.Errorf("err starting %s: %w", cmd, errStart)
	}

	errWait := c.Wait()
	if errWait != nil {
		return nil, errWait
	}

	return c, nil
}
