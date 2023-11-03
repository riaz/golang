package main

import (
	"fmt"

	human "github.com/dustin/go-humanize"
)

func main() {
	var number uint64 = 1234567
	fmt.Println("Size of the file is", human.Bytes(number))
}
