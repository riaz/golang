package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
This program finds duplicates text in a list of files passed as input

go run fileDupesInFiles.go files/sampleFile.txt files/sampleFile2.txt
*/

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// read the lines from standard input instead here, as we pass the standard input
		countLines(os.Stdin, counts) // note: map counts is passed by reference
	} else {
		for _, fname := range files {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, freq := range counts {
		if freq > 1 {
			fmt.Printf("%d\t%s\n", freq, line)
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
