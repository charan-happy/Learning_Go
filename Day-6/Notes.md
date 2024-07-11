[Pointers](#pointers)

[Panic and Recover](#panic-and-recover)

[Type Conversion](#type-conversion)

[Type Assertioin](#type-assertion)

[Type Casting](#type-casting)

## Pointers
- pointers are variables that store memory addresses of other variables
- Pointers are useful for improving for improving performance in memory-intensive tasks or for modifying variables in other functions

**Declaring and Initializing Pointers**
- Declare a pointer using * symbol followed by the data type
`var countPtr *int`
- Assign the memory address of a variable to a pointer using the & symbol
```go
var countPtr *int
count := 42
countPtr = &count
```
**Accessing and Modifying values using Pointers**
-  Access the value stored at a memory address using the * symbol before the pointer variable
`value := *countPtr // value is now 42`
- Modify the value stored at a memory address by assigning a new value using the * symbol before the pointer variable
`*countPtr = 43 // count is now 43`

**Pointers in Functions**
- pass pointers as function parameters to modify the original variable
```go
func increment(value *int) {
    *value++
}
count := 1
increment(&count) // count is now 2
```
- Return pointers from functions to share the memory address of a variable
```go
func newCounter() *int {
    counter := 0
    return &counter
}
counterPtr := newcounter() // counterPtr points to the memory address of counter
```
## Deferred Function calls
- Deferred function calls allows us to delay the execution of a function until the surrounding function returns.
- Deferred functions are useful for handling cleanup tasks, such as closing files or network connections

**Deferring functions**
- Use the defer keyword before a function call to defer its execution

```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
// output
// Hello
// World
```
### Deferred Functions and Resources
- Use defer to ensure resources are properly released, even if an error occurs during execution

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    defer file.Close()

    if err != nil {
        return err
    }
    // process file here....
}
```

### Deferred Function Call Order
- Deferred functions are called in the reverse order they were deferred (Last in,First Out or LIFO)

```go
func main() {
    defer fmt.Println("First)
    defer fmt.Println("Second")
    defer fmt.Println("Third")
}
/*
output:

Third
Second
First
*/
```

## Panic and Recover

- Panic is a mechanism to halt the normal execution of a program when an unrecoverable error occurs
- Recover is a mechanism to regain control of the program after a panic and potentially resume normal execution

- use panic built-in function to trigger a panic
- panic is typically used when an error is unrecoverable, and the program cannot proceed

```go
func main() {
    fmt.Println("Starting")
    panic("An unrecoverable error occured")
    fmt.Println("Ending") // This line will never be executed
}
```
- use recover built-in function within a deferred function to catch a panic and potentially resume normal execution

```go
func safeDivide(a, b int) int {
    defer func() {
    if r := recover(); r != nil {
    fmt.Println("Recovered from panic:", r)
    }
    }()
    return a / b
}
func main() {
fmt.Println("Starting")
result := safeDivide(10, 0)
fmt.Println("Result:", result)
fmt.Println("Ending")
}
/*
output 

Starting
Recovered from panic: runtime error: integer divide by zero
Result: 0
Endingput :
*/

```
- Recover should only be called within a deferred function, as it regains control only after a panic occurs

### panic and recover best practices
- Avoid using panic and recover as a general error-handling mechanism. Instead, use error values and error-handling techniques.
- Use panic and recover for truly exceptional situations, such as cases where the program cannot proceed due to a severe error.

## Type Conversion
- It is the process of converting a value from one datatype to another
- In Go, explicit type conversion is required when you want to convert a value from one type to another


## Type Assertion

## Type casting



