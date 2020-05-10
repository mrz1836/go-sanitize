# go-sanitize
> Simple library of sanitation methods for data sanitation and reduction

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-sanitize.svg?logo=github&style=flat)](https://github.com/mrz1836/go-sanitize/releases)
[![Build Status](https://travis-ci.com/mrz1836/go-sanitize.svg?branch=master)](https://travis-ci.com/mrz1836/go-sanitize)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-sanitize?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-sanitize)
[![codecov](https://codecov.io/gh/mrz1836/go-sanitize/branch/master/graph/badge.svg)](https://codecov.io/gh/mrz1836/go-sanitize)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-sanitize)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat)](https://mrz1818.com/?tab=tips&af=go-sanitize)

<br/>

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

<br/>

## Installation

**go-sanitize** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-sanitize
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-sanitize?tab=doc)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-sanitize?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-sanitize?tab=doc)

<details>
<summary><strong><code>Library Deployment</code></strong></summary>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                            Runs lint, test and vet
bench                          Run all benchmarks in the Go application
clean                          Remove previous builds and any test cache data
clean-mods                     Remove all the Go mod cache
coverage                       Shows the test coverage
godocs                         Sync the latest tag with GoDocs
help                           Show all make commands available
lint                           Run the Go lint application
release                        Full production release (creates release in Github)
release-test                   Full production test release (everything except deploy)
release-snap                   Test the full release (build binaries)
tag                            Generate a new tag and push (IE: tag version=0.0.0)
tag-remove                     Remove a tag if found (IE: tag-remove version=0.0.0)
tag-update                     Update an existing tag to current commit (IE: tag-update version=0.0.0)
test                           Runs vet, lint and ALL tests
test-short                     Runs vet, lint and tests (excludes integration tests)
update                         Update all project dependencies
update-releaser                Update the goreleaser application
vet                            Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](sanitize_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-sanitize) and uses [Go version 1.14.x](https://golang.org/doc/go1.14). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```shell script
make test
```

<br/>

## Benchmarks
Run the Go [benchmarks](sanitize_test.go):
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

<br/>

## Usage
- View the [examples](sanitize_test.go)
- View the [benchmarks](sanitize_test.go)
- View the [tests](sanitize_test.go)

Basic implementation:
```go
package main

import (
    "fmt"
    
    "github.com/mrz1836/go-sanitize"
)

func main() {

	// Execute and print
	fmt.Println(sanitize.IPAddress("  ##!192.168.0.1!##  "))

	// Output: 192.168.0.1
}
```

<br/>

## Maintainers

| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

<br/>

## Contributing
View the [contributing guidelines](CONTRIBUTING.md) and please follow the [code of conduct](CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&af=go-sanitize) to ensure this journey continues indefinitely! :rocket:

<br/>

## License

![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat)
