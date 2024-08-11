package main

import "fmt"

func main() {
    // declaring an array named manufacturers that stores the details (id and checked) of manufacturers
    manufacturers := [5]string{"Samsung", "Motorola", "Apple", "Sony", "Toshiba"}

    // USE of for loop
    fmt.Println("1. Available Manufacturers using Simple for loop:")
    // iterates over the manufacturers array to display the id of all the manufacturers
    for index := 0; index < len(manufacturers); index++ {
        fmt.Println("Index--> ", index, "Manufacturer at the index-->", manufacturers[index])
    }

    // USE OF for...of loop
    fmt.Println("2. Available Manufacturers using: range keyword of for loop:")
    // iterates over the manufacturers array to display all the manufacturer details
    for _, value := range manufacturers {
        fmt.Println(value)
    }

    //Print the third element of the manufacturers array using index
    fmt.Println("3. Third element of the array: ",manufacturers[2])
    
    //Replace the last element of the array with "Google"
    manufacturers[len(manufacturers)-1]="Google"
    fmt.Println("4. Replacing 'Toshiba' with 'Google'")
    fmt.Println(manufacturers)
}

