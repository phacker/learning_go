package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello cruel, curel, world!")
}

func main() {
	// creating a http.Server ourselves gives us more control over the server, like timeouts, TLS, and logger.
	s := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(helloWorld),
	}
	log.Fatal(s.ListenAndServe())
}
