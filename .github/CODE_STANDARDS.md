# Code Standards

This library follows cutting-edge Go development practices. The primary authority for workflows and conventions is [AGENTS.md](./AGENTS.md). If guidance here conflicts with that file, defer to **AGENTS.md**.

## Reference Material

We rely on these official resources:

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
- [Go examples](https://golang.org/pkg/testing/#hdr-Examples)
- [Go tests](https://golang.org/pkg/testing/)
- [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)
- [gofmt](https://golang.org/cmd/gofmt/)
- [golangci-lint](https://golangci-lint.run/)
- [Go Report Card](https://goreportcard.com/)

## AGENTS.md

All coding standards, commit conventions and pull request requirements are defined in [AGENTS.md](./AGENTS.md). Review it before contributing.

## Effective Go

Consult the [Effective Go](https://golang.org/doc/effective_go.html) documentation for idiomatic patterns and language conventions.

## golangci-lint

We run [golangci-lint](https://golangci-lint.run/usage/quick-start) to enforce style and complexity rules. Active linters live in [`.golangci.json`](../.golangci.json).

Install on macOS:
```shell
brew install golangci-lint
```

Install on Linux or Windows:
```shell
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
golangci-lint --version
```

## Documentation

All code should be documented for consumption with `godoc` and include concise examples and function comments.
