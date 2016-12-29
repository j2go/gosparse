package gosparse

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func CreateByTruncate(path string, size int64) error {
	if f, err := os.Create(path); err != nil {
		return err
	} else {
		f.Close()
	}
	if err := os.Truncate(path, size); err != nil {
		return err
	}
	if _, err := os.Stat(path); err != nil {
		os.Remove(path)
		return err
	}
	return nil
}

func CreateByDD(path string, mSize int) error {
	seek := fmt.Sprintf("seek=%d", mSize)

	out, err := exec.Command("dd", "if=/dev/null", "of="+path, "bs=1M", seek).CombinedOutput()
	if err != nil {
		return err
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		command := fmt.Sprintf("dd if=/dev/null of=%s bs=1M %s", path, seek)
		status := exitErr.Sys().(syscall.WaitStatus)
		exitCode := status.ExitStatus()
		return fmt.Errorf("command %q exits with %d: %s", command, exitCode, out)
	}
	return nil
}
