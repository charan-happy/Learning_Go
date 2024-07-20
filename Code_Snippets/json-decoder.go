package main
import (
	"encoding/json"
	"fmt"
	"os"
)
type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
func main() {
	file, err := os.Open("person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var person Person
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("JSON data successfully decoded: %+v\n", person)
}