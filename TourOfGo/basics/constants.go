package main

import "fmt"

/*
 decalared like variables but with a const prefix
*/

const Pi = 3.14

func main() {
	const World = "World"
	fmt.Println("Hello", "World")
	fmt.Println("Happy", "Pi", "Day")

	const True = true
	fmt.Println("Go rules?", True)
}
