package main

import (
	"fmt"
	"math"
)

/**
 imports can be written as
 import (
	 "os"
	"fmt"
)

or individually

import "os"
import "fmt"
import "math"

*/

// go run imports.go
func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
