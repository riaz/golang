package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
This program prints the duplicate lines in a file
This program reads a file from cli and gets the duplicate lines
*/

func main() {
	counts := make(map[string]int) // string -> int

	if len(os.Args) < 2 {
		fmt.Println("Missing parameters, provide file name")
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Can't read file: ", os.Args[1])
		panic(err)
	}
	defer file.Close()

	// reading a file from the input
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for line, freq := range counts {
		if freq > 1 {
			fmt.Printf("%d\t%s\n", freq, line)
		}
	}
}
