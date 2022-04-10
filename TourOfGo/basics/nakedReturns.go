package main

import "fmt"

/**
In go we have named returns aka naked returns, where we define the variables
as part of the func return declaration instead of part of the func.

This is recommended for short functions, and it harm readability
*/

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// go run nakedReturns.go
func main() {
	fmt.Println(split(17))
}
