# Code Standards

This project uses the following code standards and specifications from:
- [effective go](https://golang.org/doc/effective_go.html)
- [go tests](https://golang.org/pkg/testing/)
- [go examples](https://golang.org/pkg/testing/#hdr-Examples)
- [go benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
- [gofmt](https://golang.org/cmd/gofmt/)
- [golint](https://github.com/golang/lint)
- [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc)
- [vet](https://golang.org/cmd/vet/)
- [report card](https://goreportcard.com/)

### *effective go* standards
View the [effective go](https://golang.org/doc/effective_go.html) standards documentation.

### *golint* specifications
The package [golint](https://github.com/golang/lint) differs from [gofmt](https://golang.org/cmd/gofmt/). The package [gofmt](https://golang.org/cmd/gofmt/) formats Go source code, whereas [golint](https://github.com/golang/lint) prints out style mistakes. The package [golint](https://github.com/golang/lint) differs from [vet](https://golang.org/cmd/vet/). The package [vet](https://golang.org/cmd/vet/) is concerned with correctness, whereas [golint](https://github.com/golang/lint) is concerned with coding style. The package [golint](https://github.com/golang/lint) is in use at Google, and it seeks to match the accepted style of the open source [Go project](https://golang.org/).

How to install [golint](https://github.com/golang/lint):
```bash
$ go get -u golang.org/x/lint/golint
$ cd ../go-sanitize
$ golint
```

### *go vet* specifications
[Vet](https://golang.org/cmd/vet/) examines Go source code and reports suspicious constructs. [Vet](https://golang.org/cmd/vet/) uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.

How to run [vet](https://golang.org/cmd/vet/):
```bash
$ cd ../go-sanitize
$ go vet -v
```

### *godoc* specifications
All code is written with documentation in mind. Follow the best practices with naming, examples and function descriptions.