
[packages](#packages)  
[Imports](#imports)  
[Exports](#exports)  

# packages
- Packages are a way to organize and reuse code in Go programs.

**creating a package**
- Declare a package by using the package keyword followed by the package name at the beginning of a Go file

```go
// File: mypackage/mypackage.go
package mypackage
func MyFunction() {
// ...
}
```
- package names should be lowercase and descriptive

# Imports
- Imports allows you to use code from other packages in your program
- use import keyword to include external package in your program

```go
// File: main.go
package main
import  {
    "fmt"
    "mypackage"
}
func main() {
    mypackage.MyFunction()
    fmt.Println("Hello, World!")
}
```
- Use the package name followed by a dot and the item's name to access items from imported packages


# Exports
- Exports determine which items in a package are acccessible from other packages

- Export items from a package by starting their names with Uppercase

```go
// File: mypackage/mypackage.go
package mypackage
// Exported function (accessible from other pacakges)
func MyFunction() {
    // ...
}
// unexported function (not accessible from other packages)
func myPrivateFunction() {
    // ...
}
```
- only exported items can be accessed from other packages.

**Aliasing Imports**
- use aliasing to avoid naming conflicts or make imported package names more readable

```go
// File: main.go
package main
import (
    "fmt"
    mypkg "mypackage"
)
func main() {
    mypkg.MyFunction()
    fmt.Println("Hello, World!")
}
```

- use the alias followed by a dot and the item's name to access items from aliased packages.
