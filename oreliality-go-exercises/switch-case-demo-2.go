package main
import "fmt"
/*
switch expression {
case conditionA:
	// statements to be executed for condition A
case conditionB:
	// statements to be executed for condition B
default:
	// statements to be executed for default case (when cases mentioned above are not met)
*/
func main(){
var plan string = "pre"
switch plan {
case "prepaid":
	fmt.Println("plan: prepaid")
	fmt.Println("Please purchase the required service before availing it")
case "postpaid":
	fmt.Println("plan: postpaid")
	fmt.Println("please subscribe the required service and pay at the end of the month")
default:
	fmt.Println("plan: Undefined")

}
