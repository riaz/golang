package main

import "fmt"

func main() {

	messages := make(chan string)

	// send ping to messages in the same thread would lead to a deadlock.
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
