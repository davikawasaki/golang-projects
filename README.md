# GoLang Projects

Projects for Go learning and evolution.

## Table of Contents

1\. CLI project

2\. Unit testing

3\. CSV parser

4\. HTTP Module

## CLI Project

```bash
$ cd stringparse-cli
```

### Building

```bash
$ go build .
```

### Executing

```bash
$ go build .
$ go run .
$ ./stringparse
```

## Unit testing

Reference from [Medium blog post](https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318).

```bash
$ cd unitests/greetings
```

### Dependencies

- [Aurora](https://github.com/logrusorgru/aurora)

- [Gotest](https://github.com/rakyll/gotest)

- [ANSI Color Texts](github.com/fatih/color)

### Building

```bash
$ go build .
```

### Executing

```bash
$ go build .
$ go run .
$ ./greetings
```

### Testing

```bash
$ go test -v
$ $HOME/go/src/github.com/rakyll/gotest/gotest -v 
```

Testing a specific test:

```bash
$ go test -v -run TestHelloValidArg
$ $HOME/go/src/github.com/rakyll/gotest/gotest -v -run TestHelloValidArg
```

### Testing coverage

```bash
$ go test -v -cover
$ $HOME/go/src/github.com/rakyll/gotest/gotest -v -cover
```

Adding a profile to output information about coverage info in a file:

```bash
$ mkdir coverage
$ go test -v -coverprofile=coverage/cover.txt
$ $HOME/go/src/github.com/rakyll/gotest/gotest -v -coverprofile=coverage/cover.txt
```

Analyzing coverage information of a test:

```bash
$ go tool cover -html=coverage/cover.txt -o coverage/cover.html
```

### Testing multiple packages inside module with coverage

```bash
$ go test ./... -v -cover
```

### Outputing binary from specific package test

```bash
$ go test -c ./transform
```