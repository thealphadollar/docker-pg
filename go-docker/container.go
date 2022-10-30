//go:build linux
// +build linux

package main

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		must(errors.New("wrong argument"))
	}
}

// run is an implementation similar to the `docker run` command.
func run() {
	syscall.Sethostname([]byte("container"))

	// cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		GidMappingsEnableSetgroups: true,
	}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
