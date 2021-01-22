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
