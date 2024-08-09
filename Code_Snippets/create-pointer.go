package main

import (
	"fmt"
)

func main() {

	//Creating a pointer
	empName := "Rachel Green!"
	var ptrEmp = &empName
	fmt.Printf("Type of ptrEmp is %T\n", ptrEmp)
	fmt.Println("Address of empName is", &empName)

	//Zero value of a pointer is nil
	age := 25
	var ptrAge *int
	if ptrAge == nil {
		fmt.Println("\nage is", ptrAge)
		ptrAge = &age
		fmt.Println("ptrAge after initialization is", ptrAge)
	}

	//dereferencing a pointer
	salary := 25500
	ptrSalary := &salary
	fmt.Println("\nAddress of salary is", ptrSalary)
	fmt.Println("Value of salary is", *ptrSalary)

	//changing the value pointed using dereference
	*ptrSalary++
	fmt.Println("\nNew value of salary is", salary)
}
