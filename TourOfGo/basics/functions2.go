package main

import "fmt"

/**
In functions.go, we had two two parameters of the same type, and
declared them individually (int),  we may just add the type for the last varible
and implicitly it would imply that type for all the preceeding variables
*/

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(20, 30))
}
