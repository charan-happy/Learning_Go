package main
import "fmt"
func main(){
for count := 0; count <= 5; count++ {
    	for flag := 0; flag <= count; flag++ {
            	fmt.Print("*")
        }
        fmt.Println()
}

}
