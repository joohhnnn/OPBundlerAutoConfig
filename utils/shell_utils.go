package utils

import (
	"os/exec"
)

// RunShellCommand 执行shell命令
func RunShellCommand(command string) (string, error) {
	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
