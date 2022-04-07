package main

import (
	"fmt"
	"math"
)

/*
 	In go its also possible to initialize a variable withiin a if scope, followed
	 by the if condition, ideally around that variable
*/

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
