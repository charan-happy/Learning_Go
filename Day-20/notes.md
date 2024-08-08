# soft object oriented programming

## index
* Introduction
- Explanation
    - structs
    - Interfaces
    - composition
- Examples
    - Structs and Methods
    - Interfaces and polymorphism
    - Composition
    - Pseudo inheritance
- Advantages and Disadvantages
    - Advantages
    - Disadvantages
- Real-world Applications
- conclusion


## Introduction
- Go or Golang is a statistically typed and compiled programming language developed by google for simplicity and performance
- Go is primarily known for its focus on concurrency and ease of use in developing developing scalable systems, it also offers unique approach to OOP(object oriented programming) 

- Unlike traditional languages like Java, c++ and Python etc. Go does not have classes, inheritance or rigid class hierarchy . Instead go adopts a  more flexible and practical approach to oop using structs, interfaces and composition.

i.e; In Go, Object oriented programming is achieved through a combination of structs, interfaces and composition. 

**structs**
- In Go, structs are used to define custom data types that group together related data. They serve as a primary built in blocks for modeling objects and encapsulate data. Structs can also have methods,which are functions that operate on the data associated with a struct instance.

```go
type person struct {
    Name string
    Age int
}
func (p Person) Greet() {
    fmt.Println("Hello, my name is %s, and i am %d years old. \n", p.Name, p.Age)
}
```

**Interfaces**
- Interfaces in Go are a way to define a contract or a set of behaviors that a type must implement. They provide a level of abstraction and enable polymorphism, allowing different types to be be treated as if they were the same type and based on their behavior. Interfaces are defined using 'interface' keyword and a list of method signature that a type must implement to satisfy the interface.

```go
type Greeter interface {
    Greet()
}
```
- In the above example, any type that implements the Greet method can be considered a Greeter.

**composition**
- Go favors composition over inheritance,as it provides more flexible way to build complex structures and share behavior among different types. 
- Composition is achieved by embedding one struct type within another, effectively reusing the fields and methods of the embedded type.

```go
type employee struct {
    Person
    Position string
}
func (e Employee) Work(){
    fmt.Printf("%s is working as a %s. \n", e.Name, e.Position)
}
```
- In the above example, the Employee struct embeds the Person struct, inheriting its fields and methods.

By using structs, interfaces and composition. Go offers a practical and flexible approach to object oriented programming. allowing developers to model complex systems, encapsulate data and behavior, and promote code reuse and maintainability

## Examples

EX 1: **Structs and methods**

- Suppose we want to model a simple geometric shape, such as circle. We can define a circle struct and methods to calculate its area and circumference

```go
package main
import (
    "fmt"
    "math"
)
type circle struct {
    Radius float64
}
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
func (c Circle) Circumference() float64 {
    return 2 * math.Pi * c.Radius
}
func main() {
    c := Circle{Radius: 5}
    fmt.Printf("Area: %f\n", c.Area())
    fmt.Printf("Circumference: %f\n", c.Circumference())
}
```

Ex 2: **Interfaces and Polymorphism**

Now let's consider a scenario where we have multiple geometric shapes, and we want to calculate their areas. We can define shape interface and implement it for each shape type.

```go
package main
import (
    "fmt"
    "math"
)
type Shape interface {
   Area() float64 
}
type Circle struct {
    Radius float64
}
type Rectange struct {
    Width, Height float64
}
func (c Circle) Area() float64 {
    return math.PI * c.Radius * c.Radius
}
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
func printArea(s Shape) {
    fmt.Printf("Area: %f\n", s.Area())
}
func main(){
    c := Circle{Radius: 5}
    r := Rectangle{Width: 10, Height: 20}

    printArea(c)
    printArea(r)
}
```
In these example, the print function accepts any type that implements the shape interface, allowing us to easily calculate the area of the different shape types

Ex -3: **Composition**
- Suppose we want to model a zoo, with various animals that can make sounds. We can use composition to reuse the common properties of animals and define specific behavior for each animal type

```go
package main
import (
    "fmt"
)
type Animal struct {
    Name string
}
type SoundMaker interface {
    MakeSound()
}
type Dog struct {
    Animal
}
type Cat struct {
    Animal
}
func (d Dog) MakeSound() {
    fmt.Printf("%s says: Woof!\n", d.Name)
}
func (c Cat) MakeSound() {
    fmt.Println("%s says: Meow!\n", c.Name)
}
func main() {
    dog := Dog{Animal{Name: "Buddy"}}
    cat := Cat{Animal{Name: "Whiskers"}}

    dog.MakeSound()
    cat.MakeSound()
}
```
- In this example, the Dog and Cat structs both embed the Animal struct, reusing its fields and methods. Each animal type then implements the SoundMaker interface, defining its unique behavior

Ex -4: **Pseudo Inheritance**
- In Golang, inheritance is achieved through a technique called "embedding", which involves including an instance of one struct within another struct. Embedding enables the fields and methods of the embedded struct to be directly accessed by the containing struct, effectively simulating inheritance. 

This technique favors composition over classical inheritance, allowing developers to reuse the code and share behavior among different types more flexibly.

Let's look at an example to demonstrate how inheritance can be achieved using embedded structs in golang.

Suppose we have a base struct Animal with some common properties and methods, and we wnat to create two derived structs, Dog and Cat, that inherit the properties and methods of animal

```go
package main
import "fmt"
type Animal struct {
    Name string
    Age int
}
func (a Animal) Speak() {
    fmt.Printf("%s, the %d-year-old animal, speaks \n", a.Name, a.Age)
}
type Dog struct {
    Animal
    Breed string
}
type Cat struct {
    Animal
    Color string
}
func main() {
    dog := Dob{Animal{Name: "Buddy", Age: 4}, "Golden Retriever"}
    cat := Cat{Animal{Name: "Whiskers", Age: 2}, "Black"}
    dog.Speak()
    cat.Speak()

    fmt.Printf("Buddy is a %s\n", dog.Breed)
    fmt.Printf("Whiskers is a %s cat \n", cat.Color)
}
```

In this example, The Dog and cat structs both embed the Animal struct, As a result, the Name and Age fields, as well as the Speak method, are directly accessible from instances of the Dog and Cat structs, effectively simulating inheritance.

please note that golang doesnot support classical inheritance/multiple levels of hierarchy/multiple inheritance. Embedding structs combined with interfaces can be used to model complex relationships and behavior in a more flexible and modular way.


## Advantages and Disadvantages

**Advantages**
- **Composition over inheritance:** By favoring composition over inheritance, Golang promotes code reuse and maintainability. This approach allows developers to build complex structures and share behavior among different types more flexibly.
- **Simplified Polymorphism :** Interfaces in Go provides a simple and powerful way to achieve polymorphism without the complexity of inheritance hierarchies. This makes it easier to reason about code and reduces the likelihood of design mistakes
- **Encapsulation :** Structs and methods in golang allow for data encapsulation, ensuring the internal state of object is protected and can only be accessed or modified through its defined methods

- **Clear separation of concerns :** Go's approach to OOP encourages a clear separation of concerns, making it easier to organize code and manage dependencies. This results in more maintainable and scalable applications

- **strong typing:** Golang is statistically typed language, which means that type information is available at compile time. This can help prevent run-time errors and improve code reliability.


**Disadvantages**
- **Lack of classical Inheritance:**

- **Learning Curve:**

- **Potentially verbose code:**
  - Golang’s emphasis on simplicity can sometimes result in more verbose code compared to languages that offer syntactic sugar or more advanced features, such as class-based inheritance, abstract classes, or generics. This may lead to slightly longer code in some situations.

By understanding and leveraging the strengths and trade-offs of Golang's OOP model, developers can create modular, reusable and scalable code.


## Real-World Applications

Golang's approach to object-oriented programming has been widely adopted in various real-world applications. It's focus on simplicity, performance and maintainability makes it particularly well-suited for building large-scale concurrent, and distributed systems. 

Here are some of the examples of real-world applications that can be benefit from Golang's OOP Features :

Webservices and APIs: Go's efficient handling of concurrent requests and its straightforward approach to strucuring and organizing code make it an excellent choice for building web services and APis. Go's support for interfaces and composition allows developers to create modular and easily testable components for handling different aspects of request processing, such as authentication, validation and data transformation.

Microservices : Go's lightweight and fast runtime makes it well-suited for developing microservices. The object-oriented programming features in Go, such as structs, interfaces and composition, can be used to design and implement service components with clear separation of concerns and effective encapsulation of business logic

Distributed Systems: Go's built-in support for concurrency, combined with its OOP features, make it an excellent choice for building distributed systems. Developers can model complex systems using structs, interfaces and composition, While leveraging Go's powerful concurrency primitives to build highly performant and scalable applications

Commandline tools and utilities : Go's simplicity and performance make it a popular choice for developing command-line tools and utilities. using Go's OOP features, developers can create modular and maintainable code, organizing functionality into reusable components with clear interfaces.

Networking and infrastructure tools : Go has gained popularity in the networking and infrastructure space, with projects like Docker, Kubernetes and etcd leveraging its OOP features to build robust and maintainable systems. Go's approach to OOP allows developers to create modular components that can be easily extended and reused across different projects and applications.


