package main
import "fmt"
func main(){
	fmt.Println("Below are the list of numbers that are divisible by 3 and lessthan 100")
	for i := 1; i <= 100; i++{
	if i % 3 == 0 {
		fmt.Println(i)
	}else {
		i++
	}

}
}


