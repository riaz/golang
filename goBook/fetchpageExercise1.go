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

		written, err := io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s: %v\n", url, err)
			continue
		}

		status := resp.Status

		resp.Body.Close()

		fmt.Printf("written %d bytes and returned %v \n", written, status)

	}
}
