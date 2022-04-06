package main

import "fmt"

/*
We dont need to explicitly mention type, if we are initializing the variable
and it will happen implicitly
*/

var x, y = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(x, y, c, python, java)
}
