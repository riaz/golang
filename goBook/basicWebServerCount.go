package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

/*
This program keeps a ledger of each type of http request call made , in a concurrent setting.
*/

var mu sync.Mutex // this helps in reporting the accurate count
var cnt int
var runMode = "N" // this is the default mode of the program

func main() {

	hasDebugFlag(&runMode)

	// this is ugly coding can be done better
	modeMsg := "normal"
	if runMode == "D" {
		modeMsg = "debug"
	}

	fmt.Printf("Running the program in %s mode\n", modeMsg)

	http.HandleFunc("/", counter)
	http.HandleFunc("/count", count)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()

	if runMode == "D" {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		// getting form data if any passed
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form [%q] = %q\n", k, v)
		}
	}
	cnt++

	mu.Unlock()
	fmt.Fprintf(w, "URL Path : %q\n", r.URL)
}

func count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", cnt)
	mu.Unlock()
}

// the purpose of this function is to check if any of the argument to the function says
// --debug if, that is the case the request has additional information for the user.
func hasDebugFlag(mode *string) {
	const debugFlag = "--debug"
	for _, flag := range os.Args[1:] {
		if flag == debugFlag {
			*mode = "D"
		}
	}
}
