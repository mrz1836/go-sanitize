*go-sanitize* implements a simple library of various sanitation methods for data transformation.

## Installation

*go-sanitize* requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```
$ go get -u github.com/mrz1836/go-sanitize
```

### This package adheres to *go-lint* specifications
The package [golint](https://github.com/golang/lint) differs from [gofmt](https://golang.org/cmd/gofmt/).
The package [gofmt](https://golang.org/cmd/gofmt/) formats Go source code, whereas [golint](https://github.com/golang/lint) prints out style mistakes.

The package [golint](https://github.com/golang/lint) differs from [govet](https://golang.org/cmd/vet/).
The package [govet](https://golang.org/cmd/vet/) is concerned with correctness, whereas [golint](https://github.com/golang/lint) is concerned with coding style.
The package [golint](https://github.com/golang/lint) is in use at Google, and it seeks to match the accepted style of the open source [Go project](https://golang.org/).
```
$ go get -u golang.org/x/lint/golint
$ cd ~/../go-sanitize
$ golint
```

### This package adheres to *go-vet* specifications
[Vet](https://golang.org/cmd/vet/) examines Go source code and reports suspicious constructs, such as Printf calls whose arguments
do not align with the format string. [Vet](https://golang.org/cmd/vet/) uses heuristics that do not guarantee all reports are genuine problems,
but it can find errors not caught by the compilers.
```
$ cd ~/../go-sanitize
$ go vet -v
```
