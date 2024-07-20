package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Please Enter your name")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hi, %s! I'm ready to GO", name)
}
