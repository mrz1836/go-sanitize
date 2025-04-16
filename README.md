# go-sanitize
> Simple library of sanitation methods for data sanitation and reduction

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-sanitize.svg?logo=github&style=flat)](https://github.com/mrz1836/go-sanitize/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/mrz1836/go-sanitize/run-tests.yml?branch=master&logo=github&v=3)](https://github.com/mrz1836/go-sanitize/actions)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-sanitize?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-sanitize)
[![codecov](https://codecov.io/gh/mrz1836/go-sanitize/branch/master/graph/badge.svg)](https://codecov.io/gh/mrz1836/go-sanitize)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-sanitize)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat)](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-sanitize&utm_term=go-sanitize&utm_content=go-sanitize)

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
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to GitHub and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                  Runs multiple commands
clean                Remove previous builds and any test cache data
clean-mods           Remove all the Go mod cache
coverage             Shows the test coverage
godocs               Sync the latest tag with GoDocs
help                 Show this help message
install              Install the application
install-go           Install the application (Using Native Go)
lint                 Run the golangci-lint application (install if not found)
release              Full production release (creates release in Github)
release              Runs common.release then runs godocs
release-snap         Test the full release (build binaries)
release-test         Full production test release (everything except deploy)
replace-version      Replaces the version in HTML/JS (pre-deploy)
tag                  Generate a new tag and push (tag version=0.0.0)
tag-remove           Remove a tag if found (tag-remove version=0.0.0)
tag-update           Update an existing tag to current commit (tag-update version=0.0.0)
test                 Runs vet, lint and ALL tests
test-ci              Runs all tests via CI (exports coverage)
test-ci-no-race      Runs all tests via CI (no race) (exports coverage)
test-ci-short        Runs unit tests via CI (exports coverage)
test-short           Runs vet, lint and tests (excludes integration tests)
uninstall            Uninstall the application (and remove files)
update-linter        Update the golangci-lint package (macOS only)
vet                  Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/mrz1836/go-sanitize/actions) and
uses [Go version 1.18.x](https://golang.org/doc/go1.18). View the [configuration file](.github/workflows/run-tests.yml).

Run all tests (including any integration tests)
```shell script
make test
```

<br/>

## Benchmarks
Run the Go [benchmarks](sanitize_test.go):
```shell script
make bench
```

### Benchmark Results

| Benchmark                        | Iterations |   ns/op | B/op | allocs/op |
|----------------------------------|------------|--------:|-----:|----------:|
| BenchmarkAlpha                   | 1,879,712  |   635.2 |  120 |         6 |
| BenchmarkAlpha_WithSpaces        | 2,388,141  |   502.9 |   88 |         4 |
| BenchmarkAlphaNumeric            | 1,587,366  |   757.7 |  128 |         6 |
| BenchmarkAlphaNumeric_WithSpaces | 1,911,022  |   634.4 |  112 |         4 |
| BenchmarkBitcoinAddress          | 2,156,088  |   555.2 |  160 |         4 |
| BenchmarkBitcoinCashAddress      | 1,619,050  |   744.5 |  160 |         4 |
| BenchmarkCustom                  | 879,280    | 1,279.0 |  943 |        17 |
| BenchmarkDecimal                 | 2,035,514  |   595.4 |   56 |         3 |
| BenchmarkDomain                  | 2,493,144  |   473.0 |  225 |         6 |
| BenchmarkDomain_PreserveCase     | 2,879,966  |   420.9 |  209 |         5 |
| BenchmarkDomain_RemoveWww        | 1,673,802  |   719.2 |  274 |         9 |
| BenchmarkEmail                   | 2,140,860  |   560.2 |  136 |         6 |
| BenchmarkEmail_PreserveCase      | 2,634,862  |   458.5 |  112 |         5 |
| BenchmarkFirstToUpper            | 13,146,956 |    90.7 |   24 |         1 |
| BenchmarkFormalName              | 3,300,636  |   360.5 |   64 |         3 |
| BenchmarkHTML                    | 2,541,874  |   473.1 |   64 |         3 |
| BenchmarkIPAddress               | 2,895,540  |   408.4 |   80 |         5 |
| BenchmarkIPAddress_IPV6          | 1,000,000  | 1,074.0 |  225 |         6 |
| BenchmarkNumeric                 | 2,908,365  |   414.4 |   40 |         3 |
| BenchmarkPathName                | 2,348,728  |   510.1 |   64 |         3 |
| BenchmarkPunctuation             | 1,929,290  |   624.6 |  160 |         4 |
| BenchmarkScientificNotation      | 1,955,768  |   614.5 |   56 |         3 |
| BenchmarkScripts                 | 2,020,128  |   598.1 |   64 |         2 |
| BenchmarkSingleLine              | 536,860    | 2,146.0 |   96 |         4 |
| BenchmarkTime                    | 2,159,088  |   556.3 |   40 |         3 |
| BenchmarkURI                     | 2,309,937  |   518.2 |   80 |         3 |
| BenchmarkURL                     | 2,329,815  |   514.4 |   80 |         3 |
| BenchmarkXML                     | 4,123,827  |   290.3 |   56 |         3 |
| BenchmarkXSS                     | 3,485,330  |   344.7 |   40 |         2 |


<br/>

## Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## Usage
- View the [examples](examples)
- View the [benchmarks](sanitize_test.go)
- View the [tests](sanitize_test.go)

<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |

<br/>

## Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-sanitize&utm_term=go-sanitize&utm_content=go-sanitize) to ensure this journey continues indefinitely! :rocket:

<br/>

## License

[![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat)](LICENSE)
