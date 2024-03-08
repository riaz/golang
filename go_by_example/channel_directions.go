package main

import "fmt"

/*
when using a channel as a func paramater - we can specify its directionality
i.e either send or receive message
this improves the type safety of the program
*/

// read this as ping_chan channel takes in string
func ping(ping_chan chan<- string, msg string) {
	ping_chan <- msg
}

// read this as ping_chan channel takes value from a channel (str value) and sends to pong_chan
func pong(ping_chan <-chan string, pong_chan chan<- string) {
	msg := <-ping_chan
	pong_chan <- msg
}

func main() {
	ping_chan := make(chan string, 1) // this is the sender
	pong_chan := make(chan string, 1) // this is the receiver

	ping(ping_chan, "Hello World")
	pong(ping_chan, pong_chan)

	fmt.Println(<-pong_chan)

}
