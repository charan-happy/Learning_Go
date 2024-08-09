package main
import (
    "fmt"
)
type employee struct {
    firstName string
    lastName  string
    age       int
    salary    int
}
func main() {
    emp8 := &employee{
        firstName: "Sam",
        lastName:  "Anderson",
        age:       55,
        salary:    6000,
    }
    fmt.Println("First Name:", emp8.firstName)
    fmt.Println("Age:", emp8.age)
    fmt.Println("First Name:", (*emp8).firstName)
    fmt.Println("Age:", (*emp8).age)
}

