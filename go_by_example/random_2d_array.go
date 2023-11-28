package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Define a 3x3 multi-dimensional array
	var array [3][3]int

	// Populate the array with values
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			array[i][j] = i*3 + j + 1
		}
	}

	// Print the whole array
	fmt.Println("The multi-dimensional array:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println()
	}

	// Access a specific element
	fmt.Printf("\nElement at [1][1]: %d\n", array[1][1])

	// Random access
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random indices
	randomI := rand.Intn(3)
	randomJ := rand.Intn(3)

	// Access the random element
	fmt.Printf("\nRandomly accessed element at [%d][%d]: %d\n", randomI, randomJ, array[randomI][randomJ])
}
