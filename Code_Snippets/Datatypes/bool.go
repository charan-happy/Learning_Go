package main

import "fmt"

func main() {
	isPasswordValid := true
	var isAccountValid bool // default value for bool is false
	fmt.Println("isPasswordValid:", isPasswordValid, "isAccountValid", isAccountValid)
}

/*
> In go, bool is the keyword to declare boolean data type
> Boolean values doesnot support implicit or explicit type conversion
> Default value for boolean data type is false


*/
