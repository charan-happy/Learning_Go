package main
import (
    "encoding/json"
    "fmt"
)
type Address struct {
    Street string `json:"name"`
    City string `json:"city"`
}
type Person struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Address Address `json:"address"`
}
func main() {
    jsonData := []byte(`{
    "name": "charan",
    "age": 35,
    "address": {
    "street": "123 marathalli",
    "city": "Banglore"
    }
    }`)
    var person Person
    err := json.Unmarshal(jsonData, &person)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
        return
    }
    fmt.Println("Decoded nested JSON data: %+v\n", person)
}