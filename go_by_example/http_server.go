package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Create a dummy HTTP JSON response
	// Note: having interface{} is like Any
	response := map[string]interface{}{
		"message": "Hello, world!",
	}

	// Marshal the response into JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the JSON response to the HTTP response writer
	w.Write(jsonResponse)
}

// we want to use this method to return the headers of a go http call
func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, value := range headers {
			fmt.Fprintf(w, "%v : %v\n", name, value) // this is saved as part of the writer
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", headers)

	http.ListenAndServe(":8090", nil)
}
