package main

import "fmt"

func main() {
	var customerId int = 2501 // 1. Initializing at the time of declaration
	fmt.Println(customerId)
	var validity int // 2. Initializing later (useful while using loops)
	validity = 22

	var LoginMessage = "Welcome user" // 3. Initializing and declaration with type inference

	price := 100.0     // 4. Shorthand initialization and declaration
	i, j, k := 1, 2, 3 // 5.Initializing and declaration of multiple variables
	_ = 32             // 6. Using blank identifier
	fmt.Println(validity, price, LoginMessage)
	fmt.Println(i, j, k)
}
