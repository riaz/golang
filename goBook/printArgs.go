package main

import (
	"fmt"
	"os"
)

func main() {
	var sep, s string
	sep = " > "
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}
