# go-sanitize
**go-sanitize** is a simple library of sanitation methods for data sanitation and reduction

[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-sanitize)](https://golang.org/)
[![Build Status](https://travis-ci.org/mrz1836/go-sanitize.svg?branch=master)](https://travis-ci.org/mrz1836/go-sanitize)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-sanitize?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-sanitize)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-sanitize.svg?style=flat)](https://github.com/mrz1836/go-sanitize/releases)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-sanitize?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-sanitize?tab=doc)

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Installation

**go-sanitize** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/mrz1836/go-sanitize
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-sanitize).

## Examples & Tests
All unit tests and [examples](sanitize_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-sanitize) and uses [Go version 1.14.x](https://golang.org/doc/go1.14). View the [deployment configuration file](.travis.yml).
```bash
$ cd ../go-sanitize
$ go test ./... -v
```

## Benchmarks
Run the Go [benchmarks](sanitize_test.go):
```bash
$ cd ../go-sanitize
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [examples](sanitize_test.go)
- View the [benchmarks](sanitize_test.go)
- View the [tests](sanitize_test.go)

Basic implementation:
```golang
package main

import (
	"fmt"
	"github.com/mrz1836/go-sanitize"
)

func main() {

	//Execute and print
	fmt.Println(sanitize.IPAddress("  ##!192.168.0.1!##  "))

	// Output: 192.168.0.1
}
```

## Maintainers

| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-sanitize)

## License

![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat)
