Pointers

concurrency

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

