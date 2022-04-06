package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// go run functions.go
func main() {
	fmt.Println(add(4, 13))
}
