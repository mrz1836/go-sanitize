# go-sanitize
**go-sanitize** implements a simple library of sanitation methods for data sanitation and reduction. This package was inspired by processing and protecting incoming user generated content while ensuring the data will be in the correct format. This project follows Go best practices and you can view the standards and specifications at the [end of this readme](https://github.com/mrz1836/go-sanitize#adheres-to-effective-go-standards).

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-sanitize?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-sanitize)  |  [![Build Status](https://travis-ci.com/mrz1836/go-sanitize.svg?branch=master)](https://travis-ci.com/mrz1836/go-sanitize)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-sanitize.svg?style=flat)](https://github.com/mrz1836/go-sanitize/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-sanitize?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-sanitize) | [![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-sanitize) |

## Table of Contents
- [Installation](https://github.com/mrz1836/go-sanitize#installation)
- [Documentation](https://github.com/mrz1836/go-sanitize#documentation)
- [Examples & Tests](https://github.com/mrz1836/go-sanitize#examples--tests)
- [Benchmarks](https://github.com/mrz1836/go-sanitize#benchmarks)
- [Standards](https://github.com/mrz1836/go-sanitize#standards)
- [Quickstart](https://github.com/mrz1836/go-sanitize#quickstart)
- [Contributing](https://github.com/mrz1836/go-sanitize/blob/master/CONTRIBUTING.md)

## Installation

**go-sanitize** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/mrz1836/go-sanitize
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-sanitize).

## Examples & Tests
All unit tests and [examples](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go) run via [Travis CI](https://travis-ci.com/mrz1836/go-sanitize) and uses [Go version 1.11.x](https://golang.org/). View the deployment [configuration file](https://github.com/mrz1836/go-sanitize/blob/master/.travis.yml).
```bash
$ cd ../go-sanitize
$ go test ./... -v
```

## Benchmarks
Run the generic Go benchmarks:
```bash
$ cd ../go-sanitize
$ go test -bench=.
```

## Standards
- [go tests](https://golang.org/pkg/testing/)
- [go examples](https://golang.org/pkg/testing/#hdr-Examples)
- [go benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
- [gofmt](https://golang.org/cmd/gofmt/)
- [golint](https://github.com/golang/lint)
- [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc)
- [vet](https://golang.org/cmd/vet/)
- [report card](https://goreportcard.com/)

#### *effective go* standards
View the [effective go](https://golang.org/doc/effective_go.html) standards documentation.

#### *golint* specifications
The package [golint](https://github.com/golang/lint) differs from [gofmt](https://golang.org/cmd/gofmt/). The package [gofmt](https://golang.org/cmd/gofmt/) formats Go source code, whereas [golint](https://github.com/golang/lint) prints out style mistakes. The package [golint](https://github.com/golang/lint) differs from [vet](https://golang.org/cmd/vet/).
The package [vet](https://golang.org/cmd/vet/) is concerned with correctness, whereas [golint](https://github.com/golang/lint) is concerned with coding style.
The package [golint](https://github.com/golang/lint) is in use at Google, and it seeks to match the accepted style of the open source [Go project](https://golang.org/).

How to install [golint](https://github.com/golang/lint):
```bash
$ go get -u golang.org/x/lint/golint
$ cd ../go-sanitize
$ golint
```

#### *go vet* specifications
[Vet](https://golang.org/cmd/vet/) examines Go source code and reports suspicious constructs, such as Printf calls whose arguments
do not align with the format string. [Vet](https://golang.org/cmd/vet/) uses heuristics that do not guarantee all reports are genuine problems,
but it can find errors not caught by the compilers.

How to run [vet](https://golang.org/cmd/vet/):
```bash
$ cd ../go-sanitize
$ go vet -v
```

## Usage
The testable example methods are located in the [main test file](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go).
Also view the [unit tests](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go) and [benchmarks](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go) to see the other implementations.
```golang
package main

import (
	"fmt"
	"github.com/mrz1836/go-sanitize"
)

func main() {

	//Execute and print
	fmt.Println(gosanitize.IPAddress(" 192.168.0.1 "))

	// Output: 192.168.0.1
}
```

## Contributing

View the [contributing guidelines](https://github.com/mrz1836/go-sanitize/blob/master/CONTRIBUTING.md) and follow the [code of conduct](https://github.com/mrz1836/go-sanitize/blob/master/CODE_OF_CONDUCT.md).

## License

![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat)
