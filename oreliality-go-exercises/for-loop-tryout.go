package main
import "fmt"
func main(){
	package main

import "fmt"

func main() {

// variable balance declared and initialized
    balance := 600
    fmt.Println("Balance before transaction: ", balance)

// variable amount declared and initialized
    amount := 500
    fmt.Println("Amount to withdraw: ", amount)

// if-else block
    if amount <= 0 {
        fmt.Println("Withdrawal has failed as the amount is negative")
    } else if amount > balance {
        fmt.Println("Withdrawal has failed as the balance is low")
    } else {
        balance -= amount
        fmt.Println("Withdrawal has succeeded")
    }

    fmt.Println("Balance after transaction: ", balance)
}
}
