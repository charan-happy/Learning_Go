# Arrays, maps and slices

# functions

## Arrays, slices and maps

Storing multiple values in a list can be done by utilizing `arrays`. or their more flexible cousin: `slices`. A dictionary or hash type is also available, it is called  `map` in Go.

### Arrays
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

### slices

- A slice refers to an underlying array. what makes slices different from arrays is that a slice is a pointer to an array; slices are reference types, which means that if you assign one slice to another, both refer to the same underlying array.

- For instance, if a function takes a slice argument, 
### maps

- 

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

