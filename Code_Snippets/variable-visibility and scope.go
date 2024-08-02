package main

import "fmt"

var CustomerName string = "charan" // global level variable declaration --> acessible by any package and function directly

var customerContact string = "889784554" // package level  --> Accessible by all functions in this package directly

func main() {
	var customerLname string = "Taint" // block level variable declaration --> Accessible inside this function block only

	customerFullName := fetch() + customerLname
	fmt.Println("%v \nContact No. %v", customerFullName, customerContact)
}
