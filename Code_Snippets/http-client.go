package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Response:", string(body))
}