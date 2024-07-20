// Http server that listens on port 8080

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, charan!")
	})
	http.ListenAndServe(":8080", nil)
}
