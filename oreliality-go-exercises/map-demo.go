package main
import "fmt"
func main(){
//	var x map[string]int
	x := make(map[string]int) // map of type string
	x["key"] = 100
	fmt.Println(x["key"])

	y := make(map[int]int) // map of type int
	y[1] = 10
	fmt.Println(y[1])

}

