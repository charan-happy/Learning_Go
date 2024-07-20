
package main
import (
	"encoding/json"
	"fmt"
	
)
type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
func main() {
	person := Person{
		Name: "charan",
		Age: 24,
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))
}


