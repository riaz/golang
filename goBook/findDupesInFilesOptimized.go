package main

import (
	"fmt"
	"os"
	"strings"
)

/*
This is program we read the file names, and read contents of each file at once
*/

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Unable to open file %s", file)
			continue // i.e the other files may still be valid no need to panic
		}

		// Data will contain all the lines in the file
		// we will need to split it apart and read/process individual lines
		// also data is a bytes array and we type cast it using a string

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

		for line, freq := range counts {
			if freq > 1 {
				fmt.Printf("%s\t%d\n", line, freq)
			}
		}
	}

}
