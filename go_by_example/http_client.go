package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://appliedllms.com")

	// Log.fatal(err) vs panic(err) ?

	if err != nil {
		log.Fatal(err)
	}

	// TODO: How to know while files to close
	defer resp.Body.Close()

	fmt.Println("Response status", resp.Status)

	// the scanner has access to the input - it will scan the body html source line by line
	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 20; i++ { // getting the top 20 lines
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
