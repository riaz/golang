/**
This program uses range to loop over arguments instead of a for loop
*/

package main

import (
	"fmt"
	"os"
)

// Run: > go run ch1/echo2.go Hello World Testing 1 2 3
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
