/*
> In Go, we have signed and unsigned integers to hold negative and positive values respectively.
> There are different ranges integer type variable can hold like int16, int32, uint16 and so on..
> Default value for integers is 0
> In Go, int is the keyword used to create a variable of integer type
> Type conversion for int type of variables is possible
*/

package main

import "fmt"

func main() {
	var planPrice uint8 = 225 //using 8-bit unassigned int
	fmt.Println(planPrice)
	var validityDays int16 = 32767 // using 16-bit signed int
	fmt.Println(validityDays)
}
