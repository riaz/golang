package main

import (
	"fmt"
	"slices"
)

func main() {
	var x0 [3]int                                  // all will be set to 0
	var x1 = [...]int{10, 20, 30}                  // also [...] works if explictly entering all values
	var x2 = [12]int{1, 2, 3: 100, 8: 6, 10: 2, 9} // sparse where rest are 0

	fmt.Println(x0[0], x1[1], x2[3], x2[11], len(x2), len(x1), len(x0))

	// go assumes array size at compile time, hence cannot use variable for size
	// in go array of [3]int is considered a type hence cannot alter shape

	// [...] makes a array, [] makes a array

	var x4 = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x4[11])

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	//s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y)) // prints true
	fmt.Println(slices.Equal(x, z)) // prints false
	//fmt.Println(slices.Equal(x, s)) // does not compile

	// slice append
	var x3 = []int{1, 2, 3} // empty var x3 []
	x3 = append(x3, 10)
	x3 = append(x3, 20)
	x3 = append(x3, 30)
	x3 = append(x3, 4, 5, 67, 4, 3)

	fmt.Println(x3[0], x3[1], x3[2], x3[3], x4[4])

	for _, out := range x3 {
		fmt.Println(out)
	}

	// appending one slice to another
	var slice2 = []int{19, 99, 199, 1999}
	x3 = append(x3, slice2...)

	fmt.Println("Printing concatenated slices")
	for _, cc := range x3 {
		fmt.Println(cc)
	}

}
