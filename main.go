package main

import (
	"fmt"
	"log"
	"net/http"
)

type Wstr string

func (s Wstr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", Wstr("Hello world!")))
}
