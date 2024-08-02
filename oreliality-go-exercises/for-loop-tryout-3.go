package main

import "fmt"

func main() {

// variables declared using shorthand declaration method
    balance := 6000.00
    rateOfInterest := 0.10
    interest := 0.00
    deposit := 600.00

// initializing the for loop
// try declaring the above mentioned variables inside for loop and observe the changes
    for i := 1; i <= 12; i++ {
        balance += deposit
        balance -= deposit
        interest = balance * rateOfInterest
        balance += interest
        fmt.Println("The interest for the month ", i, " is ", interest)
    }
    fmt.Println("The balance at the end of the year is ", balance)
}


