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
    jsonData := []byte(`{"name": "charan", "age":24}`)
    var person Person
    err := json.Unmarshal(jsonData, &person)
    if err != nil {
        fmt.Println("Error unmarshing JSON:", err)
        return 
    }
    fmt.Println("Decoded JSON data: %+v\n", person)
}