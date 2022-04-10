package main

import "fmt"

/*
 when using := or var, the type of the varible if not explicitly specified
 is implicitly the inference type from rhs
*/

func main() {
	var val = 25
	fmt.Printf("val is of type %T\n", val)
}
