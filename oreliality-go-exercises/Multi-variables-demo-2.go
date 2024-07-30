package main
import "fmt"
// multi variable declaration
var (
	a = 5
	b = 10
	c = 15
)
func main(){
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
