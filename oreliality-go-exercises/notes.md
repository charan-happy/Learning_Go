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


