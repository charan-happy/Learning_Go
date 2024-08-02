/*

mplement a program to check whether a given number is an Armstrong number.

An Armstrong number is an n-digit number that is equal to the sum of the nth powers of its individual digits.

E.g.: 371 is an Armstrong number as 33 + 73 + 13=371

        1634 is an Armstrong number as 14 + 64 + 34+ 44=1634

 

Consider the following sample inputs:

Sample Input

Sample Output

407

407 is an Armstrong number
8207

8207 is not an Armstrong number
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
	num := input()
	//Implement your code here


}

// do not modify the code below
func input() int {
	scanner := bufio.NewScanner(os.Stdin)
	// Read the line
	scanner.Scan()
	line1 := scanner.Text()

	// Convert the string to int
	num, err := strconv.Atoi(line1)
	if err != nil {
		fmt.Println(err)
	}

	// Return the values
	return num
}

