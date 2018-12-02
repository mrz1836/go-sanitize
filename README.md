# go-sanitize
Implements a simple library of sanitation methods for data transformation. This package was inspired by protecting incoming user generated content and ensuring the data will be in the right format.

| | | | | | | |
|-|-|-|-|-|-|-|
| ![MIT](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat)  |  ![Code Size](https://img.shields.io/github/languages/code-size/mrz1836/go-sanitize.svg?style=flat) |   [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-sanitize?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-sanitize) |  [![Issues](https://img.shields.io/github/issues/mrz1836/go-sanitize.svg?style=flat)](https://github.com/mrz1836/go-sanitize/issues) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-sanitize.svg?style=flat)](https://github.com/mrz1836/go-sanitize/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-sanitize?status.svg)](https://godoc.org/github.com/mrz1836/go-sanitize) | [![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com?af=go-sanitize) |


## Installation

**go-sanitize** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```
$ go get -u github.com/mrz1836/go-sanitize
```

### Package Tests
All tests are passing on [Go v1.11.2](https://golang.org/)

### This package adheres to *effective go* standards
View the [effective go](https://golang.org/doc/effective_go.html) standards.

### This package adheres to *go-lint* specifications
The package [golint](https://github.com/golang/lint) differs from [gofmt](https://golang.org/cmd/gofmt/). The package [gofmt](https://golang.org/cmd/gofmt/) formats Go source code, whereas [golint](https://github.com/golang/lint) prints out style mistakes. The package [golint](https://github.com/golang/lint) differs from [vet](https://golang.org/cmd/vet/).
The package [vet](https://golang.org/cmd/vet/) is concerned with correctness, whereas [golint](https://github.com/golang/lint) is concerned with coding style.
The package [golint](https://github.com/golang/lint) is in use at Google, and it seeks to match the accepted style of the open source [Go project](https://golang.org/).

How to install [golint](https://github.com/golang/lint):
```
$ go get -u golang.org/x/lint/golint
$ cd ~/../go-sanitize
$ golint
```

### This package adheres to *go-vet* specifications
[Vet](https://golang.org/cmd/vet/) examines Go source code and reports suspicious constructs, such as Printf calls whose arguments
do not align with the format string. [Vet](https://golang.org/cmd/vet/) uses heuristics that do not guarantee all reports are genuine problems,
but it can find errors not caught by the compilers.

How to run [vet](https://golang.org/cmd/vet/)
```
$ cd ~/../go-sanitize
$ go vet -v
```
