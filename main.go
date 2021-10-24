package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/hello/", hello)
	m.HandleFunc("/world/", world)
	log.Fatal(http.ListenAndServe(":8080", m))
}
