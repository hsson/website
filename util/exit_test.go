package util

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestShouldExit(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		MaybeExitWithError(errors.New("Some error"))
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestShouldExit")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Errorf("process ran with err %v, want exit status 1", err)
}

func TestShouldNotExit(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		MaybeExitWithError(nil)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestShouldNotExit")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Errorf("process ran with err %v, want exit status 0", err)
	}
}
