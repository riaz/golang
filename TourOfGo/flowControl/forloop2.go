package main

import "fmt"

/* this is for block without incement or initilizer */
/* this also works like while loop*/
func main() {
	sum := 0
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)
}
