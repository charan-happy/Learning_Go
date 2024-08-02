package main
import "fmt"

func main(){
	var plan string = "undefined"
	if plan == "prepaid" {
		fmt.Println("Existing plan is prepaid")
	}else if plan=="postpaid" {
		fmt.Println("Existing plan is postpaid")
	}else {
		fmt.Println("Existing plan is undefined")
	}
}
