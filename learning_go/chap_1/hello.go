package main

import "fmt"

func main() {
	fmt.Printf("Hello, %s\n", "world!")
}

// Run : go build -o hello
// Run: go fmt ./... - format in the current directories and the sub-directories
