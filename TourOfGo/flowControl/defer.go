package main

import "fmt"

/*
	defer statement defers execution until the surrounding function returns
*/

func main() {
	defer fmt.Println("world") // executes #2
	fmt.Println("hello")       // executes #1
}
