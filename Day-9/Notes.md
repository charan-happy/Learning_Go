
# Errors and Error Handling
- Errors are a way to represent and handle abnormal conditions or failures in Go program.
- Go uses built-in error interface to manage errors, allowing for a consistent error handling approach across the language

**Understanding the Error Interface**
- The error interface is defined as:

```go
typee error interface {
    Error() string
}
```
- Any type implementing the Error() string method satisfies the error interface

## Creating and Returning Errors
- Use the errors.New() function to create new error values.

```go
import "errors"
func myFunction() (string, error) {
    // ...
    if failureCondition {
        return "", errors.New("An error occured")
    }
    // ...
    return "Success, nil"
}
```
- functions can return an error value as a second returned value, allowing callers to check for errors.

**Handling Errors**
- Check for errors by comparing the returned error value with nil


```go
result, err := myFunction()
if err !=nil {
    fmt.Println("Error:", err)
}
else {
    fmt.Println("Result:", result)
}
```
- Handle errors according to your program's requirements. Ex: log the error, display a message or retry the operation.

**Custom Error Types**
- create custom error types by defining a struct and implementing the Error() string method

```go
type MyError struct {
    Message string
}
func (e MyError) Error() string {
return e.Message
}
```
- Custom error types allows you to provide additional context or functionality

**Wrapping Errors**
- Use the fmt.Errorf() function to wrap errors with additional context

```go
func myFunction() (string, err){
    // ..
    if failureCondition {
        return "", fmt.Errorf("An error occured: %w", SomeOther Error)
    }
    // ...
    return "Success", nil
}

```
- Wrapping Errors helps in preserving the original erro while adding extra information about the failure