package main

import (
	"fmt"
	"net/http"
)

func main() {
	// BUG ReverseProxy is not working when you add extra route patterns to the microservice
	http.HandleFunc("/firsthello", hello)
	http.HandleFunc("/secondhello", test)
	http.ListenAndServe(":8002", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Method)
	// r.ParseForm() // what's going on here
	// fmt.Println(r.Form)
	fmt.Fprintf(w, "Hello test, %s!\n", r.URL.Path[1:])
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}
