package main

import (
	"fmt"
	"math"
)

const s string = "constant" // if explity not mentioned it assumes it as int

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
