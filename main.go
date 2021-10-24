package main

import (
	"fmt"
	"log"
	"net/http"
)

type Wstr string

func (s Wstr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from barFunc: %s", s)
}

// If ListenAndServe has nil as the Handler, then it uses http.DefaultServeMux.
// http.Handle and http.HandleFunc register their Handlers with http.DefaultServeMux.

func main() {
	// note that localhost:8080/bar goes here
	http.Handle("/bar", Wstr("barbarbar"))

	//  while everything else goes here.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!, from the HandleFunc")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
