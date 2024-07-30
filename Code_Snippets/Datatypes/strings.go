/*
> Strings are text types used to store values like customerName, customerContact, customerAddress etc
> In Go, strings are defined as a sequence of characters. These characters are of variable width and each character is represented by one or more bytes
> Character representation in strings uses UTF-8 Encoding(variable width encoding technique)
> Strings are immutable. This implies that strings value can be replaced by new strings, but individual string values cannot be altered
> Although strings can be concatenated
> Default value of string is nothing but blank string, Not a null string
*/
package main

import "fmt"

func main() {
	loginPass := "Welcome User! \n Login Successful" // creating a string using shorthand declaration
	var age string                                   // string using var keyword
	age = "27 years"
	fmt.Println("String 1: ", loginPass)
	loginPass[2] = "No" // throw a compilation error : cannot assign to LoginPass[2]
	fmt.Println("String 2: ", age)
}
