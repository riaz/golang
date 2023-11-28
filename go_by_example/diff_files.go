package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the first file
	file1, err := os.Open("data/file1.txt")
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}
	defer file1.Close()

	// Open the second file
	file2, err := os.Open("data/file2.txt")
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}
	defer file2.Close()

	// Create scanners for both files
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	// Line numbers
	lineNumber := 1

	// Read both files line by line
	for scanner1.Scan() && scanner2.Scan() {
		line1 := scanner1.Text()
		line2 := scanner2.Text()

		// Compare lines
		if line1 != line2 {
			fmt.Printf("Difference at line %d:\nFile1: %s\nFile2: %s\n\n", lineNumber, line1, line2)
		}
		lineNumber++
	}

	// Check for errors in scanning
	if err := scanner1.Err(); err != nil {
		panic(err)
	}
	if err := scanner2.Err(); err != nil {
		panic(err)
	}
}
