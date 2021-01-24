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

## 6. Add a test

See https://golang.org/doc/tutorial/add-a-test

In the greetings directory, create a file called greetings_test.go.

```
$ go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      github.com/malsolo/go-module/greetings  0.012s
```

## 7. Compile and install the application

See https://golang.org/doc/tutorial/compile-install

```
$ cd hello

$ go list -f '{{.Target}}'
"$GOPATH"/bin/hello
```

Follow rest of instructions to add the Go install directory to your system's shell path
or the GOBIN variable using the [go env command](https://golang.org/cmd/go/#hdr-Print_Go_environment_information): ***go env -w GOBIN=/path/to/your/bin***, for instance

```
$ go env -w GOBIN=$HOME/bin
```

Run the go install command to compile and install the package:

```
$ go install
```

You can execute it by calling the reesult of the command ***go list -f '{{.Target}}'*** or just by calling hello if the previous path is in the PATH environment variable.

See also [How to Write Go Code](https://golang.org/doc/code.html)

Always take a look to [Effective Go](https://golang.org/doc/effective_go.html)


