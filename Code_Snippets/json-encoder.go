
package main

import (
	"encoding/json"
	"fmt"
	"os"
)
type Person struct{
	Name string `json:"Name"`
	Age int `json:"age"`
}
func main() {
	person := Person{
		Name: "charan",
		Age: 24,
	}
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("Error creating file:", nil)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return 
	}
	fmt.Println("JSON data successfully written to a file")
}

/*
output:

new file named person.json will be created and data we given in this file is added there.
*/
