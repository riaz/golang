package main

import "fmt"

func BubbleSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}

func MainTest() {
	fmt.Println("Bubble sort in Go")
	unsorted_input := []int{5, 3, 4, 1, 2}
	sorted := BubbleSort(unsorted_input)
	fmt.Println(sorted)
}
