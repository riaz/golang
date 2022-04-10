package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	rand.Seed(int64(os.Getpid()))
	fmt.Println("My favorite number is", rand.Intn(10))
}
