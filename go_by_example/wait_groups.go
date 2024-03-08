package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(idx int) {
	fmt.Printf("Working %d starting\n", idx)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", idx)
}

func main() {
	var wg sync.WaitGroup

	// we will create a group of 5 go-routines
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done() // waiting for the goroutine to finish
			worker(i)
		}()
	}

	// we need to wait to prevent the main thread from existing
	wg.Wait()
}
