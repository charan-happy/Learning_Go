package main

import "fmt"

func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	return a / b
}

func main() {
	fmt.Println("Starting")
	result := safeDivide(10, 0)
	fmt.Println("Result:", result)
	fmt.Println("Ending")
}
