/*
The below tryout will give you hands-on experience with call by value and call by reference functions.

Uncomment swapRef() to try call by reference.
Uncomment swapVal() to try call by value.
(Uncomment only one function at a time.)

When you execute each function analyze the output to understand the working.
*/
package main

import "fmt"

func main() {
num1 := 100
num2 := 200
fmt.Println("Before swap num1:",num1,"\nBefore swap num2:",num2)
//swapRef(&num1,&num2)
//swapVal(num1,num2)
fmt.Println("After swap num1:",num1,"\nAfter swap num2:",num2)
}

func swapRef(num3 *int, num4 *int){

var temp int
temp=*num3
*num3=*num4
*num4=temp

}

func swapVal(num3, num4 int){

var temp int
temp=num3
num3=num4
num4=temp

}

