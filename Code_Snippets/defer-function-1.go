package main

import "fmt"

/*
func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
*/

func main() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
}
