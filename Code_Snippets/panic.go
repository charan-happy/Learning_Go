package main

import "fmt"

func main() {
	fmt.Println("starting")
	panic("An unrecoverable error occured")
	fmt.Println("Ending")
}
