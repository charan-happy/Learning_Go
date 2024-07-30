package main
import "fmt"
func main(){
var feet float64
fmt.Print("Enter the length in feet: ")
fmt.Scan(&feet)

meters := feetToMeters(feet)
fmt.Printf("%.2f feet is approximately %.2f meters. \n", feet, meters)
}

func feetToMeters(feet float64) float64 {
	return feet * 0.3048
}
