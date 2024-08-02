package main

import "fmt"

type Dictionary int // enum Directory

const (
	Contact1 Directory = iota //contact1 = 0
	Contact2                  // contact2 = 1
	Contact3                  // contact3 = 2
	Contact4                  // contact4 = 3
)

func main() {
	var NewContact Directory // Declaring a variable NewContact with type Directory
	NewContact = 2
	fmt.Println(NewContact == Contact3)
}

// use below approach to throw away default value and start indexing constants with index value = 1

/*
package main
import "fmt"
type Dictionary int  // enum dictionary
const (
_Directory = iota // throws away first value = 0
Contact1 // Contact1 =1
Contact2 // Contact2 =2 and so on...const
Contact3
Contact4
)
func main() {
var NewContact Directory // Declaring a variable NewContact with type Directory
NewContact = 2
fmt.Println(NewContact==Contact2)
}


*/
