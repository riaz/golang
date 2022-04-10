package main

import "fmt"

const (
	// creating one huge num and one small num
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big)) leads to overflow error
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
