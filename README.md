# go-sanitize
**go-sanitize** implements a simple library of sanitation methods for data sanitation and reduction.

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-sanitize?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-sanitize)  | [![Codacy Badge](https://api.codacy.com/project/badge/Grade/b11a08d5619849a0ae911d91e3bb47c7)](https://www.codacy.com/app/mrz1818/go-sanitize?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-sanitize&amp;utm_campaign=Badge_Grade) |  [![Build Status](https://travis-ci.com/mrz1836/go-sanitize.svg?branch=master)](https://travis-ci.com/mrz1836/go-sanitize)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-sanitize.svg?style=flat)](https://github.com/mrz1836/go-sanitize/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-sanitize?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-sanitize) |


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
All unit tests and [examples](sanitize_test.go) run via [Travis CI](https://travis-ci.com/mrz1836/go-sanitize) and uses [Go version 1.12.x](https://golang.org/doc/go1.12). View the [deployment configuration file](.travis.yml).
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

[@MrZ1836](https://github.com/mrz1836)

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-sanitize)

## License

![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat)
