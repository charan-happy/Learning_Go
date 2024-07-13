# Working with Json

[Introduction to JSON in Golang]()  
[Golang standard libraries for JSON]()  
[Encoder and decoder JSON data]()  
[Marshaling and unmarshaling JSON data]()  
[Handling nested JSON data and custom types]()  
[Error handling in JSON operations]()  
Best practices for working with JSON in Golang  
— Struct tags  
— Validation  
— Omitempty and Ignore  
[Differences between marshaling/unmarshaling and encoding/decoding]()  
[Working with JSON and maps in Golang]()  
[JSON for API calls: examples and use cases]()  


## Introduction to JSON in Golang
- JSON (javascript object notation) is a lightweight data interchange format that is easy to read and write. it is widely used in webdevelopment for exchange of data between client and server
- Go lang, as a modern programmming language, has a built-in support for working with json data, making it easy for developers to parse, generate and manipulate JSON data in their applications.


## Golang standard libraries for JSON
- Golang provides built-in support for working with JSON data through its standard library, specifically the `encoding/json` package. This package offers a set of functions and types for encoding and decoding JSON data, making it easy to convert between JSON and native GOlang data structures

- Here are the most commonly used types and functions in the encoding/json package:  

**json.Marshal**: This function takes a datastructure as input and returns a JSON encoded byte slice. It is used for converting GOlang datastructures into JSON format.

**json.Unmarshal**: This function takes JSON-encoded byte slice and pointer to the target datastructure as input. It decodes the JSON data into the provided target data structure

**json.Encoder**: This type provides a way to encode JSON data directly into an io.Writer stream. It is useful when working with large JSON data or when encoding JSON data to a network connection or a file

**json.Decoder**: This type provides a way to decode JSON directly from an io.Reader stream. It is useful when working with large JSON data or when decoding JSON data from a network connnection or a file.

**json.RawMessage**: This type represents raw, unprocessed JSON data. It allows you to defer thee decoding of JSON data or pass through JSON data without modification


## Encoder and decoder JSON data
- In Golang, we can use json.Encoder and json.Decoder to encode and decode JSON data directly to and from an io.writer and io.Reader stream respetively.
- This approach is particularly useful when working with large JSON data or when encoding/decoding JSON data to/from a network connection or a file

**Encoding JSON Data**
- To Encode a GOlang datastructure into JSON, you can use json.NewEncoder function, which takes an io.Writer as an argument and returns a *json.Encoder
- You can then call the Encode method on the *json.Encoder to write the JSON data to the io.Writer

```go
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
    person := Person{
        Name: "Charan",
        Age: 25,
    }
    file, err := os.Create("person.json")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    err = encoder.Encode(person)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return 
    }
    fmt.Println("JSON data successfully written to file")
}
```
- In these example, we created a Person Strut, instantiate it, and then encode it to a JSON file using json.Encoder

**Decoding JSON Data**
- To decode JSON data into a GOlang data structure, you can use the json.NewDecoder function, which takes an io.Reader as an argument and returns *json.Decoder. you can then call the Decode method on *json.Decoder to read the JSON data from io.Reader and decode it into the target data structure.

```go
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
```

- In these example, we read the JSON data from a file and decode it into a Person struct using the json.Decoder.



## Marshaling and unmarshaling JSON data

- In Golang, marshaling and unmarshaling are the processes of converting data betweeen native Golang datastructure and JSON format. These processes are performed using json.Marshal and json.Unmarshal functions from the encoding/json package

**Marshaling JSON Data**
- Marshaling is the process of converting a Golang data structure into a JSON-encoded byte slice. The json.Marshal function takes a datastructure as input and returns a JSON-encoded byte slice and an error(if any)

```go
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
```
- In these example, we create a Person struct, instantiate it, and then marshal it to JSON using the json.Marshal function

**Unmarshaling JSON Data**
- Unmarshaling is the process of converting JSON encoded byte slice into a data structure. The json.Unmarshal function takes a JSON-encoded byte slice and a pointer to the target datastructure as input. It decodes the JSON data into the provided target datastructure and return an error (if any)

```go
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
```
- In these example, we decode JSON data from a byte slice and unmarshal it into a person struct using the json.Unmarshal function.

## Handling nested JSON data and custom types
- When working with more complex JSON data structures that include nested JSON objects or custom types, you'll need to define corresponding nested GOlang structs or custom types to correctly marshal and unmarshal the data.

**Nested JSON Data**
- To handle nested JSON data, create nested structs in Golang that match the JSON structure

```go
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
```
- In this example, we define an Address struct and include it as a field within the person struct. This structure allows us to correctly unmarshal the nested JSON data.

**Custom types**
- Sometimes, you might have custom types in your JSON data that need special handling. You can implement the json. Marshaler and json. Unmarshaler interfaces for your custom types to control the marshaling and unmarshaling process.

For ex:, suppose we have a custom UserID type, and we want to marshal and unmarshal it as a string instead of its default integer representation

```go
package main
import (
 "encoding/json"
 "fmt"
 "strconv"
)
type UserID int
type User struct {
 ID UserID `json:"id"`
 Name string `json:"name"`
}
// MarshalJSON implements the json.Marshaler interface for UserID.
func (id UserID) MarshalJSON() ([]byte, error) {
 return json.Marshal(strconv.Itoa(int(id)))
}
// UnmarshalJSON implements the json.Unmarshaler interface for UserID.
func (id *UserID) UnmarshalJSON(data []byte) error {
 var str string
 if err := json.Unmarshal(data, &str); err != nil {
  return err
 }
 parsedID, err := strconv.Atoi(str)
 if err != nil {
  return err
 }
 *id = UserID(parsedID)
 return nil
}
func main() {
 jsonData := []byte(`{"id": "123", "name": "Alice"}`)
 var user User
 err := json.Unmarshal(jsonData, &user)
 if err != nil {
  fmt.Println("Error unmarshaling JSON:", err)
  return
 }
 fmt.Printf("Decoded custom type JSON data: %+v\n", user)
}
```
- In this example, we define a custom UserID type and implement the json.Marshaler and json.Unmarshaler interfaces for it. This allows us to control how the UserID type is marshaled and unmarshaled to and from JSON data

## Error handling in JSON operations
- In Golang, error handling is an essential part of working with JSON data. Most JSON encoding and decoding operations return an error value, which you should always check and handle appropriately. Let's go through some of the common error handling techniques and scenarios.

**Checking for Errors**
- When using JSON encoding and decoding functions, always check the returned error value. If an error occurs, you should handle it accordingly, such as logging the error, returning an appropriate HTTP response, or propagating the error up the call stack.

```go
jsonData := []byte(`{"name": "Charan", "age": 25}`)
var person Person
err := json.Unmarshal(jsonData, &person)
if err != nil {
    fmt.Println("Error unmarshaling JSON:", err)
    return
}
```
- In this example, we check the error returned by the json.Unmarshal() function. if an error occurs,we log it and return early from the function.

**Handling Syntax Errors**
- One of the most common types of errors that can occur during JSON decoding is syntax errors. If the JSON data is malformed or contains unexpected characters, a json.SyntaxError will be returned.

- you can use a type assertion to check if the returned error is a json.SyntaxError and handle it accordingly:

```go
jsonData := []byte(`{"name": "charan", "age": "25}`) // Missing closing quote
var person Person
err := json.Unmarshal(jsonData, &person)
if err != nil {
    if syntaxErr, ok := err.(*json.SyntaxError); ok {
        fmt.Println("JSON syntax error at offset %d: %v\n", syntaxErr.offset, syntaxErr)
    } else {
        fmt.Println("Error Unmarshaling JSON:", err)
    }
    return
    }
}
```
- In these example, we use a type assertion to check if the error is a json.SyntaxError. if it is, we print a more specific error message, including the offset where the error occured.


**Handling Unknown Fields**
- By default, GOlang's  JSON decoder ignores unknown fields in the JSON data when unmarshaling it into a struct. If you want to detect and handle unknown fields, you can use the json.Decoder with the DisallowUnknownFields() option:

```go
jsonData := []byte(`{"name": "John Doe", "age": 30, "unknownField": "value"}`)
var person Person
decoder := json.NewDecoder(bytes.NewReader(jsonData))
decoder.DisallowUnknownFields()
err := decoder.Decode(&person)
if err != nil {
 if unknownFieldErr, ok := err.(*json.UnmarshalTypeError); ok {
  fmt.Printf("Unknown field '%s' at offset %d\n", unknownFieldErr.Field, unknownFieldErr.Offset)
 } else {
  fmt.Println("Error decoding JSON:", err)
 }
 return
}
```
- In this example, we create a json.Decoder and set the DisallowUnknownFields() option. If an unknown field is encountered, a json.UnmarshalTypeError error is returned, which we handle using a type assertion


## Best practices for working with JSON in Golang
— Struct tags  
— Validation  
— Omitempty and Ignore  

## Differences between marshaling/unmarshaling and encoding/decoding

## Working with JSON and maps in Golang

## JSON for API calls: examples and use cases