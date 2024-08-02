package main

import "fmt"

func main(){
package main

import "fmt"

func main() {

//variables discount and custType declared
            var discount float32
	var custType string = "Premium" 

//switch case statement 
//Try declaring variable discount with switch case condition
	switch custType {

	case "Regular":
		discount = 5
		fmt.Println("Regular customers are eligible for 5% discount")

	case "Premium"
		discount = 10
		fmt.Println("Premium customers are eligible for 5% discount")

	default: 
		discount = 0
		fmt.Println("Guest customers are not eligible for discount")
	}

       /* switch number:=5+3; number {		
 	case 1, 5, 10:
fmt.Println("The number displayed is: ",number)
case 2, 3, 8:
fmt.Println("The number displayed is: ",number)
default: 
fmt.Println("The number displayed is: ",number)
}       */
}

}
