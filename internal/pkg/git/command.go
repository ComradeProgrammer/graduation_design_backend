package git

import (
	"os/exec"
)

// if dir is empty,use current directory
func RunCommand(dir string, task string, arg ...string) (string, error) {
	cmd := exec.Command(task, arg...)
	if dir != "" {
		cmd.Dir = dir
	}
	dataBytes, err := cmd.Output()
	if err != nil {
		return "", nil
	}
	return string(dataBytes), nil
}
