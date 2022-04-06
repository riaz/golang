package main

import (
	"fmt"
	"os"
)

// this program takes arguments and prints the same
// Sample
//go run ch1/echo1.go Hello World Testing 1 2 3
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
