# Arrays, maps and slices

# functions

## Arrays, slices and maps

Storing multiple values in a list can be done by utilizing `arrays`. or their more flexible cousin: `slices`. A dictionary or hash type is also available, it is called  `map` in Go.

### Arrays
- Arrays are fixed-size, homogeneous data structures that store elements of the same type.
- To declare an array, use the `var` keyword followed by array name, the size in square brackets[], and the element type.
`var numbers [6]int`
- You can initialize an array  during declaration by providing the elements inside curly braces {}
`var numbers = [5]int{1, 2, 3, 4, 5}`
- An array is defined by: `[n]<type>`, where 'n' is the length of the array. Assigning or indexing an element in the array is done with square brackets:

```go
var arr = [10]int
arr[0] = 42
arr[1] = 13
fmt.Println("The first element is %d\n", arr[0])
```

- Array types like `var arr = [10]int` have fixed size. the `size is part of type`. They can't grow, because then they would have different type. Also arrays are values: Assigning one array to another copies all the elements. In particular, if you pass an array to function, it will receive a copy of the array, not a pointer to it.

- To declare an array you can use the following : `var a [3]int` to initialize it to something else than zero, use a composite literal: `a := [3]int{1, 2, 3}` and this can be shortend to `a := [...]{1, 2, 3}` where Go counts the elements automatically. Note that all fields must be specified. So, if you are using multi-dimensional arrays you have to do quite some typing:

`a := [2][2]int{ [2]int{1, 2}, [2]int{3, 4} }` which is same as 

`a := [2][2]int{ [...]int{1, 2}, [...]int{3, 4} }`

when declaring arrays you always have to type something in between the square brackets, either a number of three dots (...)

- Access array elements using the index enclosed in square brackets []
`firstNumber := numbers[0] // Access the first element`

### slices
- slices are dynamic, resizable, and reference a section of an underlying array.
- To declare a slice, use the `var` keyword followed by the slice name and the element type with empty square brackets []. `var numbers []int`
- create a new slice using the make() function or by slicing an existing array or slice
```go
numbers := make([]int, 5) // create a slice with a length of 5 and a capacity of 5
existingSlice := numbers[1:4] // slices the 'numbers' slice from index 1 (inclusive) to 4 (exclusive)
```
#### slice operation
- Append elements to a slice using the append() function
```go
numbers := []int{1, 2, 3}
numbers = append(numbers, 4, 5)
```
- copy elements from one slice to another using the copy() function
```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```
- change the length or capacity of a slice using `[:n]` or `[:m:n]` syntax
```go
numbers := []int{1, 2, 3, 4, 5}
numbers = numbers[:3] // Truncate the slice to the first three elements
```
- A slice refers to an underlying array. what makes slices different from arrays is that a slice is a pointer to an array; slices are reference types, which means that if you assign one slice to another, both refer to the same underlying array.

- For instance, if a function takes a slice argument, 

#### initializing slices

- slices can be created using make() function, which takes the element type, the length, and an optional capacity

`numbers := make([]int, 5, 10)`--> creates a slice with a length of 5 and capacity of 10

- slices can also be initialized by providing the elements inside curly braces {}

`var numbers = []int{1, 2, 3, 4, 5}`

- Access slice elements using the index enclosed in square brackets []
`FirstNumber := numbers[0] // Access the first element`

#### Working with Arrays and Slices
- Iterate over arrays or slices using the for loop with range keyword.

```go
for index, value := range numbers {
  fmt.Println("numbers[%d] = %d\n", index, value)
}
```
- Modify array or slice elements by assigning a new value to the specified index

`numbers[0] = 42 // set the first element to 42`

- Find the length of an array or a slice using the len() function.
`length := len(numbers)`
- for slices, find the capacity using the cap() function
`capacity := cap(numbers)`

### maps

- Maps are unordered collection of key-value pairs. keys must be unique, and values can be of any data type
- Maps are useful for tasks like counting occurences or storing configuration data.

#### Initializing Maps

- Declare an empty map using make() function
`var ages = make(map[string]int)`
- initialize a map with key-value pairs during declaration
```go
var ages = map[string]int{
  "Charan": 25,
  "Naga": 24,
}
```

#### Working with Maps
```go
// Access a map's value by it's key
age := ages["Charan"]
// Add or modify key-value pairs in a map
ages["Cherry"] = 23
ages["Charan"] = 28
// Remove a key-value pair from a map using the delete() function
delete(ages, "Charan")
// Check if a key exists in a map
age, ok := ages["Charan"]
if ok {
  fmt.Println("Charan's age is", age)
} else {
  fmt.Println("Alice's age is not in the map")
}
// iterate over key-value pairs in a map using the for loop with the range keyword
for name, age := range ages {
  fmt.Println("%s is %d years old\n", name, age)
}
// length of a map using the len() function
count := len(ages)
```
## Functions

- functions are blocks of code that can be defined to perform a specific task and can be called by their name. Functions promote code reusability and help in organizing your code in a modular way.

**Defining Function**  
- Functions are defined using the `func` keyword, followed by the function name, a list of input parameters enclosed in parenthesis, the return type(s), and a block of code enclosed in curly braces {}.

```go
// a simple function that adds two integers and returns their sum:

func add(a int, b int) int {
return a + b
}
```

- if two or more consecutive input parameters have the same type, you can specify the type once at the end

```go
func add(a, b int) int {
return a + b
}
```

**calling function**

- To call a function, simply use its name followed by the required arguments enclosed in parenthesis:

```go
sum := add(5, 3)
fmt.Println("The sum of 5 and 3 is", sum)
```

**Multiple return values**

- Functions in Go can return multiple values. This is particularly useful when you want to return both the result and an error status. To specify multiple return types, enclose them in parenthesis:

```go
func divide(a, b float64) (float64, error) {
  if b == 0 {
  return 0, fmt.Errorf
}
return a/b , nil
}
```

- In this example, the divide function returns a float64 value representing the result and an error value indicating any errors that occured during the division. if the division is successful, the error is set to nil, which signifies that no error occured.

- When calling a function with multiple return types, you need to assign each return value to a variable, or use the _(blank identifier) if you don't need a particular return value:

**variadic functions**
- functions that can accept a variable number of arguments of the same type. To declare a variadic function, you use `...(ellipsis)` notation before the type of the last parameter. This indicates that the function can accept multiple arguments of that type.

```go
// variadic function that calculate the sum of an arbitrary number of integers:

func sum(nums ...int) int {
total := 0
for _, num := range nums {
total += num
}
return total
}
```

- In this example, the nums parameter is a variadic parameter of type int. Inside the function, nums is treated as a slice of int values, so you can use a for range loop to iterate over the arguments and calculate their sum.

- To call a variadic function, simply pass the required arguments separated by commas:

```go
result := sum(1, 2, 3, 4, 5)
fmt.Println("The sum of the numbers is", result)
```

we can also pass a slice to a variadic function by using the ... notation after the slice variable:

```go
numbers := []int{1, 2, 3, 4, 5}
result := sum(numbers...)
fmt.Println("The sum of the numbers is", result)
```

- Variadic functions are useful when you want to create a function that can handle an arbitrary number (based on chance rather than being planned) of arguments of the same type. This flexibility makes it easier to write reusable and clean code

Note:  
If a function has other parameters in addition to the variadic parameter, the variadic parameter must be the last one. for ex: Consider a function that requires a string and a variable number of integers:

```go
func processStringAndNumbers(prefix string, nums ...int) {
fmt.Println("prefix:", prefix)
fmt.Println("Numbers:")
for _, num := range nums {
fmt.Println("%d ", num)
}
fmt.Println()
}
```

To call this function, provide the string argument followed by the variadic integer arguments:

`processStringAndNumbers("My numbers:", 1, 2, 3, 4, 5)`

In short, variadic functions allows you to create flexible functions that can handle a varialbe number of arguments of the same type.

**Anonymous function**

In Golang, an anonymous function is a function that is defined without a named identifier. Anonymous functions are useful when you want to create an ad-hoc function that you don't need to reuse elsewhere in your code.

```go
func () {
fmt.Println("I'm an anonymous function")
} ()
```

In these example, we define a function without a name and immediately call it. The () at the end of the function body invokes the function.

**closures**

- Closures are special case of anonymous functions. A closure is an anonymous function that references variables from outside its own body. The closure then carries references with it wherever it goes.

```go
func main() {
message := "hello, world of go.."
sayHello := func() {
fmt.Println(message)
}
sayHello()
}
```

In these example, sayHello variable holds the closure that captures the message variable. Even though message is defined outside the function, the function still has access to it. This is key feature of closures.

In summary, anonymous functions and closures in go are powerful tools in go that enable you to use functions as values, pass functions as arguments to other functions, and define functions within functions to capture and hide state. They are especially useful when dealing with higher-order functions (functions that take other functions as arguments or return them as results)


[Day-4](https://github.com/charan-happy/Learning_Go/edit/main/Day-4/)
