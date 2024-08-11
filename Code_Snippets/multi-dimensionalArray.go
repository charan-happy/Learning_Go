/*
Consider the bill of purchase by three customers for four products. If the total amount exceeds 1000 then the customer is eligible for 15% discount. Else the customer has to pay extra 5 as tax amount.
*/

package main

import "fmt"

func main() {

    const noOfCustomers int=3
    const noOfItems int=4

   //initializing a multi-dimensional array
    costOfPurchase := [noOfCustomers][noOfItems]int{{540, 280, 270, 100},{250, 600, 256, 178},{115, 280, 123, 540}}

    fmt.Println(costOfPurchase)

    finalAmount:=0
    discount:=15

    //traversing the array using for loop
    for i := 0; i < noOfCustomers; i++ {
        sum:=0
        for j := 0; j < noOfItems; j++{
            sum+=costOfPurchase[i][j]
        }
        fmt.Print("Cost of purchase : ",sum,"...")
        if sum>=1000 {
            finalAmount=sum-(sum*discount/100)
            fmt.Println(" You are eligible for 15% discount!! Final Amount : ", finalAmount)
        }else{
            finalAmount=sum+5
            fmt.Println(" You have to pay extra tax amount of 5! Final Amount : ", finalAmount)
        }

    }
}

