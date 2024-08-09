/*
Implement the logic to swap two numbers using pointers as parameters in a function. Refer to the sample output for the same:

Before SWAP: 10 20

After SWAP: 20 10
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(x *int, y *int) {
	// Implement your code here


}
func main() {
	num1, num2 := input()
	//Implement your code here

}

// do not modify the code below
func input() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	// Read the line
	scanner.Scan()
	line1 := scanner.Text()

	// Split the line by comma
	data := strings.Split(line1, ",")

	// Convert the string to int
	num1, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Println(err)
	}

	// Convert the string to int
	num2, err := strconv.Atoi(data[1])
	if err != nil {
		fmt.Println(err)
	}

	// Return the values
	return num1, num2
}

