/*
when we reinitialize a variable in differnet scope, the innermost initialization takes the highest priority. This concept is called "shadowing"
*/

package main

import "fmt"

var customerLoginId string = "xeiow" // declared at package level
var validity int = 177

func main() {
	fmt.Println(customerLoginId)         //xeiow
	var customerLoginId string = "xeios" // variable declared at block level
	validity = 100
	// reinitializing will shadow away previous value
	fmt.Println(customerLoginId)
	var customerLoginId = 28 // compilation error
	parse()
	fmt.Println(customerLoginId)
}

func parse() {
	customerLoginId = "28"
	fmt.Println(customerLoginId)
}
