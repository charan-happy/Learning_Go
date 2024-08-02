/*
> In GO, float data type stores decimal values in single or double precesion format
> Default, float value is 0. The default type for floating point is float64
> In Go, float32 or float64 are the keywords used to create a variable of float type.
> Type conversion is possible with float
*/

package main

import "fmt"

func main() {
	var averageSales float32 = 10
	quarterlyAverage := 20.1 // type of quarterly is float64
	fmt.Println("averageSales type %T, quarterlyAverage type %T", averageSales, quarterlyAverage)
}
