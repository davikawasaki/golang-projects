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

```bash
$ cd unitests/greetings
```

### Dependencies

- [Aurora](github.com/logrusorgru/aurora)

- [Gotest](github.com/rakyll/gotest)

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