package main

import (
	"fmt"
	"time"
)

/*
	Synchronizing different co-routined using a channel
	When woriking with multiple goroutines, its better to use WaitGroup
*/

func worker(done chan bool) {

	fmt.Println("working with worker 1...")
	time.Sleep(time.Second)
	fmt.Println("done working 1")

	done <- true
}

func worker2(done chan bool) {

	fmt.Println("working with worker 2...")
	time.Sleep(time.Second)
	fmt.Println("done working 2")

	done <- true
}

func main() {
	// Note: there needs to be mapping between goroutine call and done

	done := make(chan bool, 1) // size 1

	go worker(done)
	go worker2(done)
	go worker(done)

	<-done
	<-done
	<-done
}
