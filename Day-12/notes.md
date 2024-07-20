Working with HTTP Servers :

[Introduction to HTTP and APIs in Golang](#introduction-to-http-and-apis-in-golang)  
[Using Golang’s standard libraries for HTTP servers and clients](#using-golangs-standard-libraries-for-http-servers-and-clients)  
[Understanding multiplexers in Golang](#understanding-multiplexers-in-golang)  
[Creating API endpoints and managing routing](#creating-api-endpoints-and-managing-routing)  
[Handling HTTP requests and responses](#handling-http-requests-and-responses)  
[Error handling in API development](#error-handling-in-api-development)  
[Middleware and authentication](#middleware-and-authentication)  
[Context package in Golang and its use in API development](#context-package-in-golang-and-its-use-in-api-development)  
[Best practices for creating, testing, and deploying APIs](#best-practices-for-creating-testing-and-deploying-apis)  
[Popular frameworks and routers in Golang](#popular-frameworks-and-routers-in-golang)

## Introduction to HTTP and APIs in Golang
- HTTP(Hypertext Transfer Protocol) is a widely used application-layer protocol that facilitates communication between clients and servers on the internet. In Golang, we can create HTTP servers and clients to build APIs (Application Programming Interfaces) that allow different applications to communicate and exchange data

- APIs are essential in modern software development, enabling developers to create reusable and modular software components. Golang, with its robust standard library and excellent performance, is an excellent choice for building APIs

## Using Golang’s standard libraries for HTTP servers and clients
- Golang's standard library includes the `net/http` package, which provides all the necessary functionality to create HTTP Servers and clients. This package makes it easy to build APIs without relying on external dependencies.

**HTTP Server**
- To create an HTTP Server in Golang, you can use the `http.ListenAndServe()` function. This function takes two parameters: the server's address(including the port number) and the handler that defines how the server responds to incoming request. The handler can be an instance of the `http.Handler` interface or a function with the signature `func(http.ResponseWriter, *http.Request)`

Example of an HTTP Server that listens on port 8080 :

```go
package main
import (
    "fmt"
    "net/http"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, world!")
    })
    http.ListenAndServe(":8080", nil)
}
```

**Http Client:**
- the net/http package also provides the `http.Client` type for creating HTTP clients. The `http.Client` allows you to send HTTP Request to other servers and handle their responses. you can use the http.Get(), http.Post(), and other methods for different HTTP methods.

Example of an HTTP client making a GET request :

```go
package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
)
func main() {
    resp, err := http.Get("http://example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Response:", string(body))
}
```
## Understanding multiplexers in Golang
- A Multiplexer, often referred to as "mux" or "router" is responsible for routing incoming HTTP requests to the appropriate handler based on the request data like HTTP methods and URL paths. Golang's standard library provides a default multiplexer called `http.ServeMux` which is an implementation of the `http.Handler` interface
- `http.ServeMux` allows you to register patterns for the URL paths and associate them with specific handler functions. It automatically matches incoming requests to the most specific registered pattern and calls the corresponding handler

Example of using `http.ServeMux`:

```go
package main 
import (
    "fmt"
    "net/http"
)
func main() {
    mux := http.NewServeMux()
    mux.HandleFun("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, World!")
    })
    mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Greetings from your custom API!")
    })
    http.ListenAndServe(":8080", mux)
}
```

- In these example we have created a new instance of http.ServeMux and registered two routes:/ and /greet. The server listens on port 8080 and routes incoming requests to the appropriate handler function

- While the default http.ServeMux is suitable for simple usecases, many golang developers prefer using third-party routers and multiplexers for more advanced features and improved performance. popular routers include Gorilla Mux, httprouter, and Echo.


## Creating API endpoints and managing routing
- To create API endpoint and manage routing using Golang's standard library, you'll use the `http.Handle` and `http.Handlefunc` functions to register your API routes and their corresponding handlers. These handlers are responsible for processing incoming HTTP requests and sending back the appropriate response.

Step-by-step guide to create API endpoints and managing routing :

```go
//Import required packages
package main 
import (
    "fmt"
    "net/http"
)
// Defining handler functions
func greetHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello from the greet endpoint!")
}
func userHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "Listing users..")
}
// Register routes and handlers
func main() {
    http.HandleFunc("/greet", greetHandler)
    http.HandleFunc("/users", usersHandler)
    // start the HTTP server
    http.ListenAndServe(":8082", nil)
}
```

Following above steps, we've created a basic HTTP server with two API endpoints and managed their routing using Golang's standard library. this approach works for simple usecases, but for more complex applications with advanced routing needs, consider using third-party routers and multiplexers.


## Handling HTTP requests and responses

- Handling HTTP requests and responses in Golang is done through the `net/http` package. This package provides the tools to read incoming requests, create responses, and manage the communication between the client and the server

**Reading incoming requests:**
- To read the incoming HTTP request, you can access its properties and methods through the *http.Request parameter in the handler function:


```go
func handler(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    url := r.url
    header := r.Header
    // ... read other request properties and perform actions
}
```

**Parsing query parameters and form data:**
- you can parse query parameters and form data using `r.URL.Query()` and `r.ParseForm()` methods:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    queryParameters := r.URL.Query()
    r.ParseForm()
    formData := r.Form
    // .. process query parameters and form data
```
**Reading the request body:**
- To read the request body, use the `io/ioutil` package to read the `r.Body`:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        // handle error
    }
    // .. process the request body
}
```
**Setting response headers:**
- you can set the response headers using the `http.ResponseWriter's Header()` method:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Custom-Header", "custom-value")
    // .. set other response headers
}
```
**Writing the response body:**
- To write the response body, you can use the `io package` or the `fmt.Fprint` function:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, this is response body!")
}
```
**setting the response status code:**
- Set the response status code using the http.ResponseWriter's `WriteHeader()`method :

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    // .. write the response body and set other headers
}
```
By following above steps we can handle incoming HTTP Requests and create appropriate responses using the `net/http` package in golang.



