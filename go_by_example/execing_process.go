package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
	we need to exec a process to completely replace a process with another process
	note: go requires an absolute path to exec
*/

func main() {

	binary, lookErr := exec.LookPath("ls")

	if lookErr != nil {
		log.Fatal(lookErr)
	}

	args := []string{"-a", "-h", "-l"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
