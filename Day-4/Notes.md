# Structs

# Interfaces

## Structs
- Structs or Structures are custom data types that group together variables of different datatypes
- Structures are useful for organizing and representing complex data.

**Defining and Initializing Structures**
- Defining a structure using the type keyword followed by the structure name, the struct keyword, and the fields with their respective data types enclosed in curly braces {}.

```go
type Person struct {
    Name string
    Age int
    Address string
}
```
- Initialize a structure by providing values for its fields enclosed in curly braces {}

```go
charan := Person{
    Name: "Charan",
    Age: 25,
    Address: "253 Bengaluru",
}
```
- Alternatively, you can use the new keyword to create a pointer to a new structure with zero values for its fields.
`var CharanPtr *Person = new(Person)`

**Accessing and Modifying Structure Fields**
- Access structure fields using dot . notation
`name := charan.Name`
- Modify the structure fields using dot . notation
`charan.Age = 25`

**Structures in Functions**
- Pass structures as function parameters
```go
func describe(person Person) {
    fmt.Println(%s is %d years old and lives at %s\n", person.Name, person.Age, person.Address)
}
describe(charan)
```
**Return structures from functions**
```go
func newPerson(name string, age int, address string) Person {
    return Person{Name: name, Age: age, Address: address}
}
Cherry := newPerson("cherry", 26, "46 Bengaluru")
```

**Adding Methods to Structures**
- Define a method for a structure using a receiver argument
```go
func (p Person) describe() {
    fmt.Println(%s is %d years old and lives at %s\n", p.Name, p.Age, p.Address)
}
```
- call a method using dot . notation
`charan.describe()`
- Modify structure fields using a pointer receiver
```go
func (p *Person) celebrateBirthday() {
    p.Age++
}
Charan.celebrateBirthday() // Charan.Age is now 32
```
## Interfaces
- Interfaces define a set of methods that a type must implement.
- Interfaces allow for more flexible and modular code by enabling polymorphism

**Defining Interfaces**
- Define an interface using the type keyword followed by the interface name , the interface keyword, and the method signatures enclosed in curly braces {}

```go
type speaker interface {
    Speak() string
}
```
**Implementing Interfaces**
- A type implements an interface by providing concrete implementations for all methods defined in the interface

```go
type Dog struct {
    Name string
}
func (d Dog) Speak() string {
    return "Woof!"
}
```
- A variable of an interface type can hold any value that implements the interface

```go
var mySpeaker Speaker
mySpeaker = Dog{Name: "Buddy"}
```

**Using Interfaces in Functions**
- pass interface types as function parameters to accept any type that implements the interface

```go
func introduce(s Speaker) {
    fmt.Println(s.Speak())
}
introduce(mySpeaker)
```

**The Empty Interface**
- The empty interface (interface{}) has no methods, so all types implement it by default.

- The empty interface can be used when the actual type is unknown or when function should accept any type.
```go
func printAnything(a interface{}) {
    fmt.Println(a)
}
printAnything("Hello, world!")
printAnything(42)
```

[Day-5](https://github.com/charan-happy/Learning_Go/edit/main/Day-5/)