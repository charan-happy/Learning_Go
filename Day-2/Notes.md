# Variables,types and keyword
 ## Operators and built-infunction
 ## Gokeywords 
 ## Controlstructures 




 ## Variables, types and keyword
 
 Declaring and initializing variables:
- Unlike other programming languages, in Go, you can declare a variable using the `var` keyword. followed by variable name and its types:

  `var i int`
  
  This declares a variable named `i` of type int. You can assign a value to a variable using the `=` operator

  `i = 42`

  Alternatively, you can declare and initialize a variable in a single line:

  `var i int = 42`

  Go also allows us to use the short variable declaration syntax, using `:=` which infers the variable's type based on the value. `i := 42`

  Note: Short variable declaration syntax can only be used within a function.

> Multiple **var** declarations may also be grouped, **const** and **import** also allow this.

```
var (
x int
b bool
)
```

> Multiple variables of the same type can also be declared on a single line: `var i, j int` makes x and y both `int` variables. you can also make use of parallel assignment: `a, b := 20, 16`

> A special name for a variable is `_(underscore`. Any value assigned to it is discarded. `a, b := 22, 25` -> Here we only assign the integer value of 25 to b and discarded the value 22.

> Declared but otherwise unused  variables are a compiler error in GO. the following code generates this error: below `i` is declared and not used

```go
package main
func main() {
var i int
}

```
Basic Datatypes in Go :

1. **Boolean types:**
- A Boolean type represents the set of Boolean truth values denoted by the predeclared constants `true` and `false`. The boolean type is `bool`

  usecase: You can perform logical operations on booleans using the && (AND), || (OR), and ! (NOT) operators

2. **Numerical Types:**
- Go has several basic data types, including:

- **int and uint**: Signed and Unsigned integers with sizes of 8,, 16, 32, or 64 bits (e.g; int8, uint8, int16, etc). The int and uint types have a size that depends on your system architecture(32 or 64 bits).
- int8,int16, int32, int64: The range depends on the size, e.g., int8 ranges from -128 to 127, and int64 ranges from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
- uint8, uint16, uint32, uint64: The range is from 0 to the maximum positive value for their size, e.g., uint8 ranges from 0 to 255

usecase : Integers can be used with arithmetic operators such as +, -, *, /, and %.

```
Difference between int and unit

In Go, the difference between `int` and `uint` lies in their **signedness**:

1. **`int` (Signed Integer):** The `int` type represents signed integers. It can hold both positive and negative values. The size of `int` depends on the platform: 32 bits on a 32-bit system and 64 bits on a 64-bit system⁵. Use `int` when dealing with whole numbers, regardless of sign.

2. **`uint` (Unsigned Integer):** The `uint` type represents unsigned integers. It only contains positive numbers (including zero). You might use `uint` for specific scenarios, such as bitwise operations or when working with bit patterns¹. However, be cautious not to enforce or suggest that a number must be positive using unsigned types¹.

Remember, Go's standard library uses `uint` in various places, but `int` is more commonly used due to its simplicity and natural behavior¹. If you encounter situations where you need to handle non-negative values specifically, consider using `uint`. Otherwise, stick with `int` for most integer-related tasks. 🚀

```
2. **float32 and float64**: Floating-point numbers with single (32-bit) and double (64-bit) precision.

float32: This is a single-precision floating-point number that occupies 32 bits of memory. It can represent approximately 7 decimal digits of precision.  
float64: This is a double-precision floating-point number that occupies 64 bits of memory. It can represent approximately 15 decimal digits of precision and is the default type for floating-point literals

usecase : Floating-point numbers can be used with arithmetic operators such as +, -, *, and /

 ```go
 package main
 func main() {
 var a int ←Generic integer type
 var b int32  ←32 bits integer type
 a = 15 
 b = a + a  ←Illegal mixing of these types
 b = b + 5  ←5 is a (typeless) constant, so this is OK
 }
```
4. **string**: A sequence of characters, enclosed in double quotes. The string type in GO is immutable, which means once created it cannot be changed. However, we can create new strings by concatenating or modifying existing strings.
   
   `s := "hello, welcome to go lang learning journey`  
 strings can be concatenated using the `+` operator

```go
greeting := "hello"
name := "Golang"
message := greeting + ", " + name + "!"
```

To work with individual characters of a string, you can access them using an index:
```go
s := "Hello, Golang"
firstChar := s[0] // firstChar will be a byte representing the character 'H'
```

Note: Strings in Go are UTF-8 encoded by default.

**Type Inference**

   Automatic deduction of a datatype from a given expression is in programming is refereed to as `type inference`

  - Go can infer the type of a variable based on the value assigned to it. This is done using the short variable declaration syntax with the := operator:

```go
x := 45 // x is infered to be of type int
y := 3.14 // y is inferred to be of type float64
z := true // z is inferred to be of type bool
s := "Hello, Go" // s is inferred to be of type string
```
Note: once a variable has been declared with a specific type, its type cannot be changed.

**constants**
- Constants are unchanging values that are declared using const keyword. Like variables, constants can have a type, but the type can often be inferred based on value:

```go
const Pi float64 = 3.14159
const Greeting = "Hello, Go" //Type inferred as string
```

we can also use `iota` keyword to create a sequence of constants with incrementing integer values, typically within a const block:

```go
const (
A = iota // 0
B = iota // 1
C = iota // 2

```

Complex Numbers:

- Go has native support for complex numbers. if you use them you need a  variable of the type `complex`. If you need to specify the number of bits you have `complex32` and `complex64` for 32 and 64 bits.

```go
var c complex = 5+5i; fmt.Printf("value is: %v", c")
// output
// 5+5i
```
 ## Operators and built-in function

 Operators:
 
Go supports various operators for arithmetic, comparison, logical operations, and more. Here are some common ones:

**Arithmetic Operators:**
+: Addition
-: Subtraction
*: Multiplication
/: Division
%: Remainder (modulo)
**Comparison Operators:**
==: Equal to
!=: Not equal to
<: Less than
>: Greater than
<=: Less than or equal to
>=: Greater than or equal to
**Logical Operators:**
&&: Logical AND
||: Logical OR
!: Logical NOT
**Other Operators:**
&: Address of
*: Pointer dereference
++: Increment
--: Decrement

```go
package main

import "fmt"

func main() {
    // Arithmetic operators
    p := 34
    q := 20
    result1 := p + q
    fmt.Printf("Result of p + q = %d\n", result1)

    result2 := p - q
    fmt.Printf("Result of p - q = %d\n", result2)

    result3 := p * q
    fmt.Printf("Result of p * q = %d\n", result3)

    result4 := p / q
    fmt.Printf("Result of p / q = %d\n", result4)

    result5 := p % q
    fmt.Printf("Result of p %% q = %d\n", result5)

    // Relational operators
    a := 5
    b := 5
    fmt.Printf("Is a equal to b? %v\n", a == b)
    fmt.Printf("Is a not equal to b? %v\n", a != b)
    fmt.Printf("Is a greater than b? %v\n", a > b)
    fmt.Printf("Is a less than b? %v\n", a < b)
    fmt.Printf("Is a greater than or equal to b? %v\n", a >= b)
    fmt.Printf("Is a less than or equal to b? %v\n", a <= b)

    // Logical operators
    x := true
    y := false
    fmt.Printf("x AND y = %v\n", x && y)
    fmt.Printf("x OR y = %v\n", x || y)
    fmt.Printf("NOT x = %v\n", !x)
}

```
Built-in Functions:

Go provides several built-in functions that you can use directly in your programs. Some important ones include:
len: Returns the length of strings, arrays, slices, maps, and channels.  
cap: Returns the capacity (maximum storage) of slices and maps.  
make: Creates a new, initialized slice, map, or channel.  
new: Allocates memory for a value (but doesn’t initialize it).  
append: Appends elements to a slice.  
copy: Copies elements from one slice to another.  
panic and recover: Used for error handling and panic recovery.  
panic and panicln are used for an exception mechanism  
print and println arelowlevelprinting functions that can be used without reverting to
 the fmt package. These are mainly used for debugging.
cmplx, real and imag all deal with complex numbers.

Ifyouwanttoputtwo(ormore)statements
 on one line, they must be separated with a semicolon (’;’), like so:
 `a = 5; a = a + 1`
 
 ## Gokeywords 
 
```
break
 default
case 
 defer
 else
 func
 go
 goto
 fallthrough if
 interface select
 map
 package
 range
 import return
 struct
 switch
 type
 var
```

 ## Control structures 
 - control structures in Golang allows you to control the flow of your program's execution. This includes making decisions based on conditions and repeating a block of code a certain number of times.

**if**
- The if statement is used to conditionally execute a block of code. The code within the if block is executed if the specified condition evaluates to true.

```go
x := 32
if x > 30 {
fmt.Println("x is greater than 30")
}
```

we can also include an optional initialization statement before the condition:

```go
if y := x - 10; y > 0 {
fmt.Println("y is greater than o")
}
```

if-else

- The if-else statement allows you to execute a different block of code if the specified condition evaluates to false. Here's an example:

```go
x := 5
if x > 10 {
fmt.Println("x is greater than 10")
} else {
fmt.Println("x is not greater than 10")
}
```

Else-if
- The else if statement allows you to chain multiple conditions in an if-else block. When the first if condition is false, the program checks the subsequent else if conditions in order. If one of the else if conditions is true, the corresponding block of code is executed, and the remaining else if and else blocks are skipped. If none of the conditions are met, the else block is executed
  
```go
x := 25
if x > 50 {
 fmt.Println("x is greater than 50")
} else if x > 20 {
 fmt.Println("x is greater than 20 but not greater than 50")
} else {
 fmt.Println("x is not greater than 20")
}
```

Switch 

- The switch statement is used to execute a specific block of code based on the value of an expression. It is particularly useful when you have multiple conditions to check:
  
```go
x := 5
switch x {
 case 1:
 fmt.Println("x is 1")
 case 2:
 fmt.Println("x is 2")
 case 3, 4, 5:
 fmt.Println("x is 3, 4, or 5")
 default:
 fmt.Println("x is not 1, 2, 3, 4, or 5")
}
```
- The switch statement can also be used with an optional initialization statement, and you can use it with conditions instead of values:

```go
x := 42
switch y := x % 2; {
 case y == 0:
 fmt.Println("x is even")
 case y == 1:
 fmt.Println("x is odd")
}
```

For

- The for statement is used to create loops in Go. There are three types of for loops:

For loop with an initialization statement, condition, and post statement:

```go
for i := 0; i < 5; i++ {
 fmt.Println(i)
}
```

for loop with only condition (simmilar to a while loop in other languages):

```go
i := 0
for i < 5 {
 fmt.Println(i)
 i++
}
```

infinite loop:

```
for {
 fmt.Println("This loop will run forever!")
}
```
- To break out of a loop, use the break statement. To skip the current iteration and continue with the next one, use the continue statement

**For range**

The for range loop is used to iterate over the elements of a collection, such as an array, slice, or map. When used with arrays or slices, the loop provides both the index and the value at that index. When used with maps, the loop provides the key and the value for each key-value pair

Here’s an example with a slice:
```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
 fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```
If you only need the index, you can use the _ (blank identifier) to ignore the value:
```go
for index, _ := range numbers {
 fmt.Printf("Index: %d\n", index)
}
```
Similarly, if you only need the value, you can omit the index:
```go
for _, value := range numbers {
 fmt.Printf("Value: %d\n", value)
}
```
Here’s an example with a map:
```go
ages := map[string]int{
"Alice": 30,
"Bob": 25,
"Carol": 35,
}
for key, value := range ages {
 fmt.Printf("Key: %s, Value: %d\n", key, value)
}
```
By using else if and for range in your Go programs, you can create more complex conditional structures and easily iterate over collections like arrays, slices, and maps.


Taking input from the user:

```go
package main

import (
"fmt"
"strings"
)

func main() {
fmt.Println("Please enter your name.. ")
var name string
fmt.Scanln(&name)
name = strings.TrimSpace(name)
fmt.Printf("Hi, %s !, I'm ready to GO !")
}

```
- save the program as `user-input.go`  and run as `go run user-input.go`
  
-  The fmt.Scanln method tells the computer to wait for input from the
 keyboard ending with a new line or (\n), character. This pauses the
 program, allowing the user to enter any text they want
-  To use the strings package you need to import it at the top of the
 program
- The fmt.Printf function takes a string, and using
 special printing verbs, (%s), it injects the value of name into the string.
 You do this because Go does not support string interpolation, which would
 let you take the value assigned to a variable and place it inside of a string.

- the TrimSpace function, from Go’s standard library
 strings package, on the string that you captured with fmt.Scanln.
 The strings.TrimSpace function removes any space characters,
 including new lines, from the start and end of a string.
