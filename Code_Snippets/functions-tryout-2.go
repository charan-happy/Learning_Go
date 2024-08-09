package main
 
import "fmt"
 
func main() {
nums(1,2,3,4)
}
func nums(nums ...int){
 
temp := 0
for _,i := range(nums){
temp += i
}
fmt.Println(temp)
}
