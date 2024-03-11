package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

/*
Think of this as using subprocess.Run in python
*/

func main() {

	dataCmd := exec.Command("date")

	dateOut, err := dataCmd.Output() // note: this returns a byte

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()

	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc=", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "python")

	// defining the pipes for input and output
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	// starting the actual grep command execution and work with results
	grepCmd.Start()
	grepIn.Write([]byte("python is a programming language \n python is a type of snake belonging to boa family\n I love to eat pie"))
	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut) // getting results from the output pipe
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// using a bash command
	lsCmd := exec.Command("bash", "-c", "ls", "-a -l -h")
	lsOut, err := lsCmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listing all the files in the current directory")
	fmt.Println(string(lsOut))

}
