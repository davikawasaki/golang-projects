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