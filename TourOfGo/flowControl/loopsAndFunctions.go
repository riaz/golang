package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	prev := 0.0
	for prev != z {
		//fmt.Println(z)
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	//fmt.Println(Sqrt(1), Sqrt(2), Sqrt(3), Sqrt(4), Sqrt(100), Sqrt(81))
	fmt.Println(Sqrt(81))
}
