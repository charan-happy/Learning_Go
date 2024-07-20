package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, from greet endpoint!")
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Listing users...")
}
func main() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/users", userHandler)
	http.ListenAndServe(":8082", nil)
}
