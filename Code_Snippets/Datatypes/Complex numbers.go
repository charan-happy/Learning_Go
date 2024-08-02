/*
> the complex datatype is used to store numbers of the form a+ib. Where 'a' is real part and 'i' is the imaginary unit or indeterminate
> Default value for complex number is 0+0i
> In Go, There are two types of complex numbers; complex64 and complex128. The Go's parser will understand 'i' as imaginary part in complex numbers
> There are 3 builtin functions for complex numbers. imag(), real() and complex()
> real() will return real part of the complex number and imag() will return the imaginary part of the complex number
> complex() takes in real, imaginary part as input and returns the complex numbers formed


*/

package main

import "fmt"

func main() {
	var num1 complex128 = 2 + 3i        // creating complex type test variable
	var num2 complex64 = complex(2, 10) // creating complex type variable using inbuilt function
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println("%v, %T \n", real(num1), real(num1)) //using inbuilt function to render real part
	fmt.Println("%v, %T \n", imag(num1), imag(num1)) // using inbuilt function to render imaginary part
}
