.
# Struct Composition and Embedding
- Struct composition is a way to build complex structures by combining simpler ones.
- Struct embedding allows you to include one struct within another, providing a form of code reuse and simplifying struct definitions

#### struct Composition
- Create complex structures by combining simpler ones using fields of other struct types

```go
type Address struct {
    Street string
    City string
    State string
}
type Person struct {
    Name string
    Age int
    Address Address
}
```
- Access composed fields using the dot notation

```go
person := Person{
    Name: "Charan Doe",
    Age: 24,
    Address: Address{
        Street: "123 gandhi st",
        City: "New marathalli",
        State: "NY",
    },
}
fmt.Println("Street:", person.Address.Street)
```

**Struct Embedding**
- Embed a struct within another by declaring an anonymous field of the embedded struct type.

```go
type Address struct {
    Street string
    City string
    State string
}
type Person struct {
    Name string
    Age int
    Address
}
```
- Access Embedded fields directly using dot notation

```go
person := Person{
    Name: "Charan Doe",
    Age: 24,
    Address: Address{
        Street: "435 cma st",
        City: "Bengaluru",
        State: "KA",
    },
}
fmt.Println("City:", person.City) // Direct access to the embedded field
```

## Methods and Embedded Structs
- Methods defined on the embedded structs are promoted to the embedding struct, allowing you to use them directly

```go
type Address struct {
    Street string
    City string
    State string
}

func (a Address) FullAddress() string {
  return fmt.Sprintf("%s, %s, %s, a.Street, a.City, a.State)
}

type Person struct {
Name: string
Age int
Address 
}

person := Person{
Name: "Charan",
Age: "24",
Address: Address{
    Street: "792 churnch st",
    City: "Banglore",
    State: "KA",
},
}
fmt.Println("Full Address:", person.FullAddress()) // using the method from the embedded struct
```

## Interface Embedding
- Interface embedding allows you to compose complex interface by combining simpler ones
- It promotes code reuse and simplifies the definition of interfaces that share common methods

#### Basic Interface Embedding

- Embed an interface within another by including it as an anonymous field

```go
type Reader interface {
    Read(P []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
type ReadWriter interface {
    Reader
    Writer
}
```
- An embedded interface brings all its methods into the embedding interface, allowing a single interface to represent multiple capabilities

### Implementing Embedded Interfaces

- To satisfy an embedded Interface, a type must implement all the methods from the embedded interfaces

```go
type File struct {
    // Implementation details
}
func (f *File) Read(p []byte) (n int, err error) {
    // Read Implmentation
}
func (f *File) Write(p []byte) (n int, err error) {
    // Write implmentation
}
// The file type now satisfies the ReadWriter interface
var rw ReadWriter = new(File)
```
- Implementing the embedded interfaces separately allows fo r more modular and flexible design

#### Interface Embedding and Method Overriding

- when embedding interfaces with overlapping method sets, the embedding interface can override methods

```go
type A interface {
    foo()
}
type B interface {
    A
    Foo()
}
```

- In this case, any type implementing interface B must provide an implementation for the Foo() method, which will satisfy both A and B interfaces.
