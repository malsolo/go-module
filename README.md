# Tutorial: Create a Go module

See [the tutorial from the Go website](https://golang.org/doc/tutorial/create-module)

## 1. Create a module

See https://golang.org/doc/tutorial/create-module 

```
$ cd $GOPATH/src/github.com/malsolo

$ mkdir greetings

$ cd greetings

$ go mod init github.com/malsolo/go-module/greetings
```

Create greetings.go.

## 2. Call your code from another module

See https://golang.org/doc/tutorial/call-module-code

```
$ cd ..
$ mkdir hello
$ cd hello
```

Create hello.go.

```
$ go mod init hello
go: creating new go.mod: module hello
```

Add "replace example.com/greetings => ../greetings" in the go.mod and then:

```
$ go build
go: found github.com/malsolo/go-module/greetings in github.com/malsolo/go-module/greetings v0.0.0-00010101000000-000000000000

$ ./hello
Hi, Gladys. Welcome!
```

## 3. Return and handle an error

See https://golang.org/doc/tutorial/handle-errors

Modify greetings/greetings.go.

At the command line in the hello directory, run hello.go

```
$ go run hello.go
greetings: empty name
exit status 1
```

## 4. Return a random greeting

See https://golang.org/doc/tutorial/random-greeting

Modify greetings/greetings.go

```
$ go build

$ ./hello
Great to see you, Gladys!

$ ./hello
Hi, Gladys. Welcome!

$ ./hello
Hail, Gladys! Well met!
```

## 5. Return greetings for multiple people

See https://golang.org/doc/tutorial/greetings-multiple-people

Modify greetings/greetings.go 

Modify hello/hello.go



