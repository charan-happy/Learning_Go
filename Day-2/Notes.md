
 ## Variables,types and keyword
 ## Operatorsandbuilt-infunction
 ## Gokeywords 
 ## Controlstructures 
 ## Arrays,slicesandmaps


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
