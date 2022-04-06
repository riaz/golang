/**
What is the list of arguments is huge, - concatenation would be slow
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

// Run: > go run ch1/echo3.go Hello World Testing 1 2 3
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
