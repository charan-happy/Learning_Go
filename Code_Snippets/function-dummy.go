package main

import "fmt"

func updateEmployeeName(emp string) {
	emp = "charan"
}
func main() {
	emp := "Charan Cherry"
	fmt.Println("value of emp before function call->", emp)
	updateEmployeeName(emp)
	fmt.Println("value of emp after function call ->", emp)
}
