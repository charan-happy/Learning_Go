[Pointers](#pointers)

[Panic and Recover](#panic-and-recover)

[Type Conversion](#type-conversion)/[Type Casting](#type-casting)

[Type Assertioin](#type-assertion)



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

- use the target type as a function to perform a type conversion

```go
var i int = 42
var f floate64 = float64(i)
```

- Type conversion does not change the value but may cause a loss of precesion depending on the source and target types

**Converting between numeric types**

- Convert between different numeric types, such as int, float64, and uint

```go
var i int32 = 10
var j int64 = int64(i)
var k float32 = 3.14
var l float64 = float64(k)
```
- Be aware of potential precesion loss or overflow when converting between numeric types with different sizes or ranges

**Converting between strings and byte slices**
- Convert between string and []byte types using type conversion

```go
s := "hello"
b := []byte(s) // b is now a byte slice: []byte{104, 386, 36, 222}
s2 := string(b) // s2 is now a string: "hello"
```

- converting between string and []byte allows for efficient manipulation of string data

**Converting between custom types**
- You can also convert between custom types, as long as their underlying types are the same

```go
type Celsius float64
type Fahrenheit float64
func toFahrentheit(c Celsius) Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
var tempC Celsius = 100
var tempF Fahrenheit = toFahrenheit(tempC)
```

## Type Assertion
- Type Assertion is a way to extract the underlyinig values of an interface type. It allows you to access the concrete value stored in an interface variable and perform operations specific to the underlying type

**performing Type Assertions**
- use the type assertion syntax to extract the underlying value from an interface variable

```go
var value interface{} = 42
num := value.(int)
fmt.Println("The number is:", num)
```
`Type assertion syntax: variable.(TargetType)`
- Use the two-value type assertion syntax to handle type assertion errors gracefully

```go
var value interface{} = "hello"
num, ok := value.(int)
if ok {
    fmt.Println("The number is:", num)
} else{
    fmt.Println("Value is not an int")
}
```
- The two-value type assertion returns the extracted value and a boolean indicating whether the type assertion was successful.

**Type Assertion with Switch**
- Use type switch to handle multiple possible types in a single statement

```go
func describe(value interface{}) {
    switch v := value.(type) {
        case int:
        fmt.Printf("Value is an int: %d\n", v)
        case float64:
        fmt.Printf("Value is a float64: %f\n", v)
        case string:
        fmt.Printf("value is a string: %s\n", v)
        default:
        fmt.Printf("value is of unknown type: %T\n", value)
    }
}
```
- The type keyword in a Type switch allows you to determine the underlying type of an interface value

## Type casting

- Type Casting/ type conversion, is the process of converting a value from one data type to another
- It allows you to perform operations between different data types and use values of one type where another type is expected

Basic Type casting:

- Use type casting to convert a value from one datatype to another

```go
var i int = 42
var f float64 = float64(i)
fmt.Println("Integer value:", i)
fmt.Println("Float value:", f)
```

`Type casting syntax: TargetType(value)`

**Type casting and Arithmetic Operations**
- Type casting is often necessary when performing arithmetic operations between different data types.

```go
var i int = 10
var f float64 = 3.5
sum := float64(i) + f
fmt.Println("Sum:", sum)
```
- Go lang requires explicit type casting for operations involving mixed data types

**Type casting with strings**
- Use type casting to convert between string and numeric types

```go
var i int = 42
str := strconv.Itoa(i) // convert integer to string
fmt.Println("String value:", str)
num, err := strconv.Atoi(str) // convert string to integer
if err == nil {
    fmt.Println("Integer value:", num)
}
else {
    fmt.Println("Error converting string to integer:", err)
}
```
- The strconv package provides functions for converting between strings and numeric types

