# 🚀 Code Standards

Welcome to a modern Go codebase. This library follows best-in-class practices for clarity, performance, and maintainability. Our single source of truth is [AGENTS.md](./AGENTS.md). If anything here ever contradicts it, follow **AGENTS.md**.

<br/>

## 📄 Reference Material

When in doubt, check the official docs:

* ✨ [Effective Go](https://golang.org/doc/effective_go.html)
* ⚖️ [Go Benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
* 📖 [Go Examples](https://golang.org/pkg/testing/#hdr-Examples)
* ✅ [Go Testing Guide](https://golang.org/pkg/testing/)
* 📃 [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)
* 🔧 [gofmt](https://golang.org/cmd/gofmt/)
* 📊 [golangci-lint](https://golangci-lint.run/)
* 📈 [Go Report Card](https://goreportcard.com/)

<br/>

## 🧰 AGENTS.md

Everything from naming conventions to pull request etiquette lives in [AGENTS.md](./AGENTS.md). Read it. Bookmark it. Trust it.

<br/>

## 🎓 Effective Go

We adhere to the patterns and philosophy in [Effective Go](https://golang.org/doc/effective_go.html). Stick to idiomatic code. Avoid cleverness when clarity wins.

<br/>

## 🔍 golangci-lint

We lint all the things. Our active ruleset lives in [`.golangci.json`](../.golangci.json).

### 🌎 macOS, Linux or Windows

Running `make lint` will detect if it's installed. If not, it will attempt to automatically install it for you.

```sh
make lint
```

<br/>

## 📑 Documentation

All exported code must be documented. Use `godoc`-compatible comments. If your function needs an example, include it. If it doesn’t, question if it should be exported.

<br/>

Happy coding — keep it clean, idiomatic, and readable. ✨
