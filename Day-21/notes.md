Anatomy of Go workspace :

```
.

|---bin

|---pkg

|---src
    
    ^--- github.com/foo/bar
        
        ^-- bar.go
    ```
- `$ GOPATH/bin` directory is where Go places binaries that `go install` compiles. Our Os uses `$PATH` environment variable to find binary applications that can execute without a full path. It is recommended to add this directory to our global `$PATH` variable

For ex: if we don't add `$ GOPATH/bin` to `$PATH` to execute a program from there, we would need to run : `$GOPATH/bin/myapp` 

when `$GOPATH/bin` is added to `$ PATH` we can make the same call like such : `myapp`

- The `$ GOPATH/pkg` directory is where GO stores pre-compiled object files to speed up subsequent compiling of programs. 

- The `src` directory is where all of our `.go` files, or source code, must located.

## what are packages ?
- Go code is organized in packages . A package represents all the files in a single directory on disk. One directory can contain only certain files from the same package . Packages are stored, with all user-written go source files, under `$GOPATH/src` directory.

- we can understand package resolution by importing differnt packages

- if our code lives at `$GOPATH/src/blue/red` then its package name should be `red`

- The import statement for the `red` package would be :
`import "blue/red"`

- packages that live in source code repositories like github, bitbucket have the full location of the repository as part of their import path.

ex: source code at `https://github.com/gobuffalo/buffalo` using the import path:

`import "github.com/gobuffalo/buffalo"
`
