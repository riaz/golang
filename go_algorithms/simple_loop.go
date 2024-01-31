package main

import "fmt"

func loopTest() {
	for i := 0; i <= 4; i += 1 {
		for j := 0; j <= i; j += 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
