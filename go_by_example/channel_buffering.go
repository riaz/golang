package main

import "fmt"

func main() {
	messages := make(chan string, 2) //setting up channel size to 2

	messages <- "Hello"
	messages <- "World"

	// we expect to see the messages printed in order
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
