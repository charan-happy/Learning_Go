package main
import (
    "fmt"
)
func change(age *int) {
    *age = 55
}
func modify(sls []int) {
    sls[0] = 90
}
func main() {
    //passing pointer to a function
    a1 := 58
    fmt.Println("\nvalue of a1 before function call is", a1)
    ptrA1 := &a1
    change(ptrA1)
    fmt.Println("value of a1 after function call is", a1)

    //slice as an argument to a function
    o := [3]int{89, 90, 91}
    modify(o[:])
    fmt.Println("modify3 using slice", o)
}

