package main

import "fmt"

// Run - is going to be responsible of instantiation
// and startup of go application
func Run() error {
	fmt.Println("Stating up an application")
	return nil
}

func main() {
	fmt.Println("REST API")
	// This avoids the main function from panicing
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
