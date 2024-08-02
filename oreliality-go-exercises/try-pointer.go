package main
import (
    "fmt"
)
func main() {
    //declaring and initializing variable greeting
    var greeting string = "Hello World"
    //declaring string pointer, prtGreeting 
    var prtGreeting *string
    //initialization of pointer
    prtGreeting = &greeting 
    fmt.Println("value of greeting", greeting )
    fmt.Println("address of greeting", &greeting )
    fmt.Println("value of prtGreeting", prtGreeting)
}

