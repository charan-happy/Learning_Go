/* problem statement

​​​​Implement a program to find and display the maximum number out of the given three numbers.

Consider the following sample inputs and outputs:

Sample Input

Sample Output

num1 = 21, num2 = 24, num3 = 32	
32

num1 = 12, num2 = 21, num3 = 14	
21

Sample Data
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1, num2, num3 := input()
	//Implement your code here

}

// do not modify the code below
func input() (int, int, int) {
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

	// Convert the string to int
	num3, err := strconv.Atoi(data[2])
	if err != nil {
		fmt.Println(err)
	}

	// Return the values
	return num1, num2, num3
}


