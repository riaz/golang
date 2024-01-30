package main

import (
	"command-line-arguments/Users/riaz/Projects/golang/restful-prod-api/internal/db/db.go"
	"context"
	"fmt"
)

// Run - is going to be responsible of instantiation
// and startup of go application
func Run() error {
	fmt.Println("Stating up an application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("REST API")
	// This avoids the main function from panicing
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
