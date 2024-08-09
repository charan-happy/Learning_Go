package main
 
import "fmt"
 
func main() {
num1 := 10
num2 := 20
fmt.Println(add(num1,num2))
fmt.Println(sub(num1,num2))
}
func add(num3,num4 int) int{
add := num3+num4
return add
}
 
//write a function to subtract 2 numbers
func sub(num3,num4 int) int{
	sub := num3-num4
	return sub
}
