package main

import (
	"fmt"
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// LogReqeusts prints r.URL.Path to stdout
func LogRequest() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v\n", r.URL.Path)
			f(w, r)
		}
	}
}

// Method returns StatusBadRequest unless the method matches s.
func Method(s string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != s {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}

// Chain applies middlewares in reverse order to f
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", Chain(hello, Method("GET"), LogRequest()))
	log.Fatal(http.ListenAndServe(":8080", m))
}
