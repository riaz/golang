package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// go run fetchpage.go http://www.pixls.ai

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s: %v\n", url, err)
			continue
		}

		fmt.Printf("%s\n", body)

	}
}
