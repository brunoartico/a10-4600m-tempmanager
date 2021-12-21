package powerlimits

import (
	"os/exec"
	"syscall"
)

func runWithoutWindow(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return command.Run()
}

func outputWithoutWindow(cmd string, args ...string) ([]byte, error) {
	command := exec.Command(cmd, args...)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return command.Output()
}
