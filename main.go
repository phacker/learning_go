package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("before: %v\n", r.URL.Path)
		defer log.Println("after logging")
		f(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/hello/", logging(hello))
	m.HandleFunc("/world/", logging(world))
	log.Fatal(http.ListenAndServe(":8080", m))
}
