package main
import "fmt"
func fahreinheitToCelsius(fahreinheit float64) float64 {
	celsius := (fahreinheit - 32) * 5 / 9
	return celsius
}
func main(){
	result := fahreinheitToCelsius(68)
	fmt.Println("68^oF in Celsius:", result, "^oc")
}
