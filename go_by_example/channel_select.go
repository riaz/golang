package main

import (
	"fmt"
	"time"
)

/*
	go select lets u wait on multiple channel operations.
	We can use a standard select statement to await both these values
	We will use -> time go run channel_select.go -> to verify that goroutines are running concurrently
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second) // sleep for a second
		c1 <- "one"
	}()

	go func() {
		time.Sleep(4 * time.Second) // sleep for 2 seconds
		c2 <- "two"
	}()

	// now that we have two go-routines that are running and we are also in the main
	// thread and we use select as a way to await for these values.

	// Note at each cycle we expect one of the case statements to activate
	// hence the number of case statements must match the count of go coroutines.

	for i := 0; i < 2; i++ { // the purpose of the loop is to wait for 2 cyles for both job to finish
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)
		}
	}

}
