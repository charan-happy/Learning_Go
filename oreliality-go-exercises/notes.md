`package main` This is the package declaration, and every go program must starts with it. packages are go's way of organizing and reusing code There are two types of Go programs: executables and libraries. 
- Executable applications are kind of programs that can run directly from the terminal (on windows ends with .exe) 
- Libraries are collections of code that we package together so that we can use them in other programs.

- import is the keyword how we include code from other packages to use with our program.

To know something in go we can use `godoc fmt println`

- Type help us reason about what our program is doing and catch a wide variety of common mistakes

Go types

**1.Numbers**
    Integers and floating-point numbers
#### Integers :
- Go's integer types are uint8, uint16, uint32, uint64, int8, int16, int32 and int64. uint means "unsigned integer" while int means "signed integer". Unsigned integers only contain positive numbers (or zero). 
- In addition, there two alias types: byte(which is same as uint8) and rune(which is same as int32).
- There are also three machine dependent integer types: uint, int and uintptr. They are machine dependent because their size depends on the type of architecture you are using.

Generally, if you are working with integers, you should just use the `int` type

####Floating-Point Numbers:
- Floating point numbers are the numbers that contain a decimal component(i.e; real numbers) Ex: 1.234, 123.4, 0.00001234 etc.

- Go has two floating-point types: float32 and float64 (often referred as single precesion and double precesion respectively). It also has two additional types for representing complex numbers (numbers with imaginary part): Complex64 and complex128. Mostly we stick with float64 when working with floating-point numbers

- Go supports the following standard arithmetic operations:
    > Addition(+)
    > substraction(-)
    > Multiplication(*)
    > Division(/)
    > Remainder(%)

    **2.strings**
    - a string is a sequence of characters with a definite length used to represent text. Go strings are made up of individual bytes, usually one for each character.
    - String literals can be created using double quotes "Hello, world" or backticks(``).

```
The difference between double quotes and backticks is that double-quoted strings cannot contain newlines and they allow special escape sequences. 
For ex: `\n` gets replaced with a newline and `\t` gets replaced with a tab character
```
**3.Booleans**
- A boolean value (named after George Boole) is a special 1-bit integer type used to represent `true` and `false` . Three logical operators that are used with boolean values:

```
&& and 
|| or
! not
```
**variables**
- Variables is a storage location, with a specific type and an associated name.
- In Go, variables are created first by using `var` keyword then specifying variable name(x) and the type(string) and finally assigning a value to the variable("Hello, world"). Value assignment is optional

####scope of a variables
####Naming convention
####constants
- constants are esentially variables whose values cannot be changed later. Instead of using `var` keyword, we use the `const` keyword.

**Control Structures**
- unlike other programming languages. Go, has only one for loop as control structure
- for, if and switch are the main control flow statements in go

simmilar to any other programming languages, we can implement the control flow in GO using
- if-else or if-else if-else(decesion)
- switch-case(selection)
- for (iteration)

**important aspects of switch cases in go:**
- Duplicate or multiple cases with simmilar conditions are not allowed. The compiler will throw error for the same

```go
var plan string = "Pre"
switch plan {
case "Prepaid":
	fmt.Println("Plan: Prepaid")
              fmt.Println("Please purchase the required service before availing it!")
case "Prepaid":
	fmt.Println("Plan: Postpaid")
              fmt.Println("Please subscribe the required service and pay at the end of the month!")
default:
	fmt.Println("Plan: Undefined")
}

```

- An individual case can contain multiple values for the comparision

```go
var plan string = "Prepaid"
switch plan {
case "Prepaid", "Postpaid":
	fmt.Println("Plan: Prepaid")
	fmt.Println("Plan: Postpaid") 
default:
	fmt.Println("Plan: Undefined")
}
```

- The default case can come in between other cases. there is no strict rule that it should come at the end only. 

```go
var plan string = "Prepaid"
	switch plan {
	case "Prepaid":
		fmt.Println("Plan: Prepaid")
	default:
		fmt.Println("Plan: Undefined")
	case "Postpaid":
		fmt.Println("Plan: Postpaid")
        }

```

- The Alternative syntax we discussed for if-else statements can also be implemented in switch-case as well

```go
/*
switch statement; expression{
case conditionA:
         //Statements to be executed for condition A
case conditionB:
         //Statements to be executed for condition A
//Simllarly other cases to be mentioned here
default:
       //Statements to be executed for default case (when cases mentioned above are not met)
}

*/
switch plan:= "Prepaid"; plan {
	case "Prepaid":
		fmt.Println("Plan: Prepaid")
	case "Postpaid":
		fmt.Println("Plan: Postpaid")
	default:
		fmt.Println("Plan: Undefined")
}


```
**loops in go**
- unlike all other programming languages, we have only for loop in go. It does not have while and do-while loops

```go
/*
for variable initialization; condition; updation {
       //Statements to be executed in each iteration of the loop
} 

*/
for count :=1; count <-10; count++ {
fmt.Println(" This is iteration no. %d \n", count)
}
```

> Applications of for loops in go

1. Nested for loops
```go
for count := 0; count <= 5; count++ {
    	for flag := 0; flag <= count; flag++ {
            	fmt.Print("*")
        }
        fmt.Println()
}

```

2. for loop without initialization and updatation: 

```go
count := 0		//Variable initialization is done here
for count <= 10 {        
        fmt.Println(count)
 	count += 2	//Loop updation is done here
}

```

3. For loop with multiple variable initialization

```go
for num, flag := 1, 1; flag <= 5 && num <= 5; flag, num = flag+1, num+1 { 
        	fmt.Printf("%d * %d = %d\n", num, flag, num*flag)
}

```
4. infinite for loop
```go
for {                     //Initialization, Condition and Updation of the loop are completely skipped
        fmt.Println("Let’s suffer this Infinite Loop together!")
}

```
5. String traversal using string and range :

```go
for index, element := range "Dasvidanya" {
	fmt.Println(index, " => ", string(element))
}

//Accessing indices separately of the string. Here element variable is not used
for index, _ := range "Dasvidanya" {
    fmt.Printf("%v ",index)
}
fmt.Printf("\n")
//Accessing elements separately of the string.
// Here index variable is replace with underscore “_” to ignore the indices
for _, element := range "Dasvidanya" {
    fmt.Println(string(element))
}

```

6. Range over int

```go
/*
for i := range n{
 //do something 
}

*/
package main
import "fmt"
func main() {
 var num int = 10
    for i := range num {
        fmt.Println(i)
    }
}

```
7. Break statements with for loop:

```go
for count := 1; count <= 10; count++ {
 	if count > 5 {
        break            //Loop is terminated if count>5
    }
        fmt.Printf("This is loop iteration No. %d \n", count)
}

```
8. Using Continue statements with for loop:

```go
for count := 1; count <= 10; count++ {
    if count%2 == 0 {
		continue
    }
        fmt.Printf("This is iteration No. %d. Hence, iteration No. %d is skipped. \n", count, count+1)
}

```

**Pointers in Go**

- Pointers are the variables that store the address of the variable instead of the actual value of the variable or in short
- Pointer is a variable that stores the memory address of the another variable

Declaring pointer

`var pointer_name *data_type` 

- * is the special character which is termed as **deferencing operator** used to declare pointer variable and access the value store in the address

```
why we use *string as datatype ?

if you take a pointer of the string type, then it can be assigned an address of only string data type variable, not any other type. So, above is a pointer of type string which can store only the memory addresses of string variables. We can avoid specifying the data type during the declaration of pointer, as the type of pointer variable can be determined by the compiler based on the variable
```

**Initialization of pointer**

`var greeting = "Hello world"`

Initialization of pointer with memory address of variable b: `var ptrGreeting *string = &greeting`

Different ways of Creating Pointers;

1. Pointer using data-type:

```go
package main
import (
	"fmt"
)
func main() {
	//declaring and initializing variable greeting
	var greeting string = "Hello World"
	//declaring and initializing string pointer, ptrGreeting
	var ptrGreeting *string = &greeting
	fmt.Println("value of greeting", greeting)
	fmt.Println("address of greeting", &greeting)
	fmt.Println("value of ptrGreeting", ptrGreeting)
}

```

2. Shorthand Declaration :

```go
package main
import (
    "fmt"
)
func main() {
    //declaring and initializing variable greeting
    greeting:= 20
    //declaring and initialization pointer, ptrGreeting
    ptrGreeting := &greeting
    fmt.Println("value of greeting", greeting)
    fmt.Println("address of greeting", &greeting)
    fmt.Println("value of ptrGreeting", ptrGreeting)
}

```

3. Built-in New function:

```go
package main
import (
    "fmt"
)
func main() {
    //declaring and initializing variable greeting
    var greeting string = "Hello World"
    //declaring string pointer, ptrGreeting
    ptrGreeting := new(string)
    //initialization of pointer
    ptrGreeting = &greeting
    fmt.Println("value of greeting", greeting)
    fmt.Println("address of greeting", &greeting)
    fmt.Println("value of ptrGreeting", ptrGreeting)
}

```
- Go allows us to compare two pointers having same data type for equality using == operator

```go
package main
import "fmt"
func main() {
	var num = 90
	var ptr1 = &num
	var ptr2 = &num
	if ptr1 == ptr2 {
		fmt.Println("Pointer ptr1 and ptr2 points to the same variable")
	}
}    
```



`
**Arrays**:
- An array is a numbered sequence of elements of a single type with fixed length. 
- In both arrays and strings indexing starts with '0'

**slices**:
- A slice is a segment of an array. Like arrays, slices are indexable and have a length . Unlike arrays, this length is allowed to change.
- To create a slice, we have to use the built in `make` function
- slices are always associated with some array, and although they can never be longer than the array, they can be longer than the array, they can be smaller. 

- In addition to the indexing operator, Go includes two builtin functions to assist with slices: append and copy

**append**
- - Append adds elements onto the end of slice. If there's sufficient capacity in the underlying array, the element is placed after the last element and the length is incremented. 
- However, if there is not sufficient capacity a new array is created, all of the existing elements are copied over, the new element is added onto the end, and the new slice is returned.

- slices are typically used to represent list of items, particularly when you need to access the nth item quickly - for ex, player#33 or the 10th most popular search query. 

**Maps**:
- A map is an unordered collection of key-value pairs (maps are also sometimes called as an associated arrays, hash tables or dictionaries). Maps are used to lookup a value by its associated key. 

Ex: Map in Go `var x map[string]int` --> keytype in brackets and finally the value type. We can read it as "x is a map of strings to ints"

```
> Runtime errors are errors happen when you run the program
> Compiletime errors are errors that happen when you try to compile the program
```
- maps have dynamic length, maps are not sequential, we can also delete items from a map using the built-in `delete` function: `delete(x, 1)

**functions**:
- A function (also known as procedures or a subroutine) is an independent section of code that maps zero or more input parameters to zero or more output parameters. 
- Dividing our program into smaller blocks of reusable code. These blocks are called functions . Dividing code into several functions increases the code maintainability and reduces redundancy. 

- packages are defined as a group of functions and methods.

