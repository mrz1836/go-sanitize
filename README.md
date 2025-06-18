# go-sanitize
> Lightweight Go library providing robust string sanitization and normalization utilities

<table>
  <thead>
    <tr>
      <th>CI&nbsp;/&nbsp;CD</th>
      <th>Quality&nbsp;&amp;&nbsp;Security</th>
      <th>Docs&nbsp;&amp;&nbsp;Meta</th>
      <th>Community</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td valign="top" align="left">
        <a href="https://github.com/mrz1836/go-sanitize/releases">
          <img src="https://img.shields.io/github/release-pre/mrz1836/go-sanitize?logo=github&style=flat" alt="Latest release">
        </a><br/>
        <a href="https://github.com/mrz1836/go-sanitize/actions">
          <img src="https://img.shields.io/github/actions/workflow/status/mrz1836/go-sanitize/run-tests.yml?branch=master&logo=github&style=flat" alt="Build status">
        </a><br/>
        <a href="https://github.com/mrz1836/go-sanitize/commits/master">
		  <img src="https://img.shields.io/github/last-commit/mrz1836/go-sanitize?style=flat&logo=clockify&logoColor=white" alt="Last commit">
		</a>
      </td>
      <td valign="top" align="left">
        <a href="https://goreportcard.com/report/github.com/mrz1836/go-sanitize">
          <img src="https://goreportcard.com/badge/github.com/mrz1836/go-sanitize?style=flat" alt="Go Report Card">
        </a><br/>
		<a href="https://codecov.io/gh/mrz1836/go-sanitize">
          <img src="https://codecov.io/gh/mrz1836/go-sanitize/branch/master/graph/badge.svg?style=flat" alt="Code coverage">
        </a><br/>
        <a href="https://github.com/mrz1836/go-sanitize/actions">
          <img src="https://github.com/mrz1836/go-sanitize/actions/workflows/codeql-analysis.yml/badge.svg?style=flat" alt="CodeQL">
        </a><br/>
        <a href=".github/SECURITY.md">
          <img src="https://img.shields.io/badge/security-policy-blue?style=flat&logo=springsecurity&logoColor=white" alt="Security policy">
        </a><br/>
        <a href=".github/dependabot.yml">
          <img src="https://img.shields.io/badge/dependencies-automatic-blue?logo=dependabot&style=flat" alt="Dependabot">
        </a>
      </td>
      <td valign="top" align="left">
        <a href="https://golang.org/">
          <img src="https://img.shields.io/github/go-mod/go-version/mrz1836/go-sanitize?style=flat" alt="Go version">
        </a><br/>
        <a href="https://pkg.go.dev/github.com/mrz1836/go-sanitize?tab=doc">
          <img src="https://pkg.go.dev/badge/github.com/mrz1836/go-sanitize.svg?style=flat" alt="Go docs">
        </a><br/>
        <a href=".github/AGENTS.md">
          <img src="https://img.shields.io/badge/AGENTS.md-found-40b814?style=flat&logo=openai" alt="AGENTS.md rules">
        </a><br/>
        <a href="Makefile">
          <img src="https://img.shields.io/badge/Makefile-supported-brightgreen?style=flat&logo=probot&logoColor=white" alt="Makefile Supported">
        </a>
      </td>
      <td valign="top" align="left">
        <a href="https://github.com/mrz1836/go-sanitize/graphs/contributors">
          <img src="https://img.shields.io/github/contributors/mrz1836/go-sanitize?style=flat&logo=contentful&logoColor=white" alt="Contributors">
        </a><br/>
        <a href="https://github.com/sponsors/mrz1836">
          <img src="https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat" alt="Sponsor">
        </a><br/>
        <a href="https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-sanitize&utm_term=go-sanitize&utm_content=go-sanitize">
          <img src="https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat" alt="Donate Bitcoin">
        </a>
      </td>
    </tr>
  </tbody>
</table>

<br/>

## 🗂️ Table of Contents
* [Installation](#-installation)
* [Usage](#-usage)
* [Documentation](#-documentation)
* [Examples & Tests](#-examples--tests)
* [Benchmarks](#-benchmarks)
* [Code Standards](#-code-standards)
* [AI Compliance](#-ai-compliance)
* [Maintainers](#-maintainers)
* [Contributing](#-contributing)
* [License](#-license)

<br/>

## 📦 Installation

**go-sanitize** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-sanitize
```

<br/>

## 💡 Usage

Here is a basic example of how to use go-sanitize in your Go project:

```go
package main

import (
    "fmt"
    "github.com/mrz1836/go-sanitize"
)

func main() {
	// Sanitize a string to remove unwanted characters
	input := "Hello, World! @2025"
	sanitized := sanitize.AlphaNumeric(input, false) // true to keep spaces

	// Output: "Sanitized String: HelloWorld2025"
	fmt.Println("Sanitized String:", sanitized) 
}
```

- Explore additional [usage examples](examples) for practical integration patterns
- Review [benchmark results](#benchmark-results) to assess performance characteristics
- Examine the comprehensive [test suite](sanitize_test.go) for validation and coverage
- Fuzz tests [are available](sanitize_fuzz_test.go) to ensure robustness against unexpected inputs

<br/>

## 📚 Documentation

View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-sanitize?tab=doc)

<br/>

### Features
- Alpha and alphanumeric sanitization with optional spaces
- Bitcoin and Bitcoin Cash address sanitizers
- Custom regular expression helper for arbitrary patterns
- Precompiled regex sanitizer for repeated patterns
- Decimal, domain, email and IP address normalization
- HTML and XML stripping with script removal
- URI, URL and XSS sanitization

### Functions
- [`Alpha`](sanitize.go): Remove non-alphabetic characters, optionally keep spaces
- [`AlphaNumeric`](sanitize.go): Remove non-alphanumeric characters, optionally keep spaces
- [`BitcoinAddress`](sanitize.go): Filter input to valid Bitcoin address characters
- [`BitcoinCashAddress`](sanitize.go): Filter input to valid Bitcoin Cash address characters
- [`Custom`](sanitize.go): Use a custom regex to filter input _(legacy)_
- [`CustomCompiled`](sanitize.go): Use a precompiled regex to filter input **(suggested)**
- [`Decimal`](sanitize.go): Keep only decimal or float characters
- [`Domain`](sanitize.go): Sanitize domain, optionally preserving case and removing www
- [`Email`](sanitize.go): Normalize an email address
- [`FirstToUpper`](sanitize.go): Capitalize the first letter of a string
- [`FormalName`](sanitize.go): Keep only formal name characters
- [`HTML`](sanitize.go): Strip HTML tags
- [`IPAddress`](sanitize.go): Return sanitized and valid IPv4 or IPv6 address
- [`Numeric`](sanitize.go): Remove all but numeric digits
- [`PhoneNumber`](sanitize.go): Keep digits and plus signs for phone numbers
- [`PathName`](sanitize.go): Sanitize to a path-friendly name
- [`Punctuation`](sanitize.go): Allow letters, numbers and basic punctuation
- [`ScientificNotation`](sanitize.go): Keep characters valid in scientific notation
- [`Scripts`](sanitize.go): Remove scripts, iframe and object tags
- [`SingleLine`](sanitize.go): Replace line breaks and tabs with spaces
- [`Time`](sanitize.go): Keep only valid time characters
- [`URI`](sanitize.go): Keep characters allowed in a URI
- [`URL`](sanitize.go): Keep characters allowed in a URL
- [`XML`](sanitize.go): Strip XML tags
- [`XSS`](sanitize.go): Remove common XSS attack strings


### Additional Documentation & Repository Management

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

This project uses [goreleaser](https://github.com/goreleaser/goreleaser) for streamlined binary and library deployment to GitHub. To get started, install it via:

```bash
brew install goreleaser
```

The release process is defined in the [.goreleaser.yml](.goreleaser.yml) configuration file.

To generate a snapshot (non-versioned) release for testing purposes, run:

```bash
make release-snap
```

Before tagging a new version, update the release metadata in the `CITATION.cff` file:

```bash
make citation version=0.2.1
```

Then create and push a new Git tag using:

```bash
make tag version=x.y.z
```

This process ensures consistent, repeatable releases with properly versioned artifacts and citation metadata.

</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands

```bash script
make help
```

List of all current commands:

<!-- make-help-start -->
```text
all                      Runs multiple commands
citation                 Update version in CITATION.cff (citation version=X.Y.Z)
clean-mods               Remove all the Go mod cache
coverage                 Shows the test coverage
diff                     Show the git diff
generate                 Runs the go generate command in the base of the repo
godocs                   Sync the latest tag with GoDocs
govulncheck-install      Install govulncheck for vulnerability scanning
help                     Show this help message
install                  Install the application
install-go               Install the application (Using Native Go)
install-releaser         Install the GoReleaser application
lint                     Run the golangci-lint application (install if not found)
release                  Full production release (creates release in GitHub)
release-snap             Test the full release (build binaries)
release-test             Full production test release (everything except deploy)
replace-version          Replaces the version in HTML/JS (pre-deploy)
run-fuzz-tests           Runs fuzz tests for all packages
tag                      Generate a new tag and push (tag version=0.0.0)
tag-remove               Remove a tag if found (tag-remove version=0.0.0)
tag-update               Update an existing tag to current commit (tag-update version=0.0.0)
test                     Runs lint and ALL tests
test-ci                  Runs all tests via CI (exports coverage)
test-ci-no-race          Runs all tests via CI (no race) (exports coverage)
test-ci-short            Runs unit tests via CI (exports coverage)
test-no-lint             Runs just tests
test-short               Runs vet, lint and tests (excludes integration tests)
test-unit                Runs tests and outputs coverage
uninstall                Uninstall the application (and remove files)
update-linter            Update the golangci-lint package (macOS only)
vet                      Run the Go vet application
```
<!-- make-help-end -->

</details>

<details>
<summary><strong><code>GitHub Workflows</code></strong></summary>
<br/>

| Workflow Name                                                                | Description                                                                                                            |
|------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [auto-merge-on-approval.yml](.github/workflows/auto-merge-on-approval.yml)   | Automatically merges PRs after approval and all required checks, following strict rules.                               |
| [codeql-analysis.yml](.github/workflows/codeql-analysis.yml)                 | Analyzes code for security vulnerabilities using GitHub CodeQL.                                                        |
| [delete-merged-branches.yml](.github/workflows/delete-merged-branches.yml)   | Deletes feature branches after their pull requests are merged.                                                         |
| [dependabot-auto-merge.yml](.github/workflows/dependabot-auto-merge.yml)     | Automatically merges Dependabot PRs that meet all requirements.                                                        |
| [pull-request-management.yml](.github/workflows/pull-request-management.yml) | Labels PRs by branch prefix, assigns a default user if none is assigned, and welcomes new contributors with a comment. |
| [release.yml](.github/workflows/release.yml)                                 | Builds and publishes releases via GoReleaser when a semver tag is pushed.                                              |
| [run-tests.yml](.github/workflows/run-tests.yml)                             | Runs all Go tests and dependency checks on every push and pull request.                                                |
| [stale.yml](.github/workflows/stale.yml)                                     | Warns about (and optionally closes) inactive issues and PRs on a schedule or manual trigger.                           |
| [sync-labels.yml](.github/workflows/sync-labels.yml)                         | Keeps GitHub labels in sync with the declarative manifest at `.github/labels.yml`.                                     |

</details>

<details>
<summary><strong><code>Updating Dependencies</code></strong></summary>
<br/>

To update all dependencies (Go modules, linters, and related tools), run:

```bash
make update
```

This command ensures all dependencies are brought up to date in a single step, including Go modules and any tools managed by the Makefile. It is the recommended way to keep your development environment and CI in sync with the latest versions.

</details>

<br/>

## 🧪 Examples & Tests

All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/mrz1836/go-sanitize/actions) and use [Go version 1.18.x](https://go.dev/doc/go1.18). View the [configuration file](.github/workflows/run-tests.yml).

Run all tests:

```bash script
make test
```

<br/>

## ⚡ Benchmarks

Run the Go [benchmarks](sanitize_benchmark_test.go):

```bash script
make bench
```

<br/>

Performance benchmarks for the core functions in this library, executed on an Apple M1 Max (ARM64):

### Benchmark Results

| Benchmark                                             | Iterations |   ns/op | B/op | allocs/op |
|-------------------------------------------------------|------------|--------:|-----:|----------:|
| [Alpha](sanitize_benchmark_test.go)                   | 15,108,703 |    78.7 |   24 |         1 |
| [Alpha_WithSpaces](sanitize_benchmark_test.go)        | 13,972,903 |    83.2 |   24 |         1 |
| [AlphaNumeric](sanitize_benchmark_test.go)            | 10,619,542 |   112.0 |   32 |         1 |
| [AlphaNumeric_WithSpaces](sanitize_benchmark_test.go) | 10,005,721 |   118.9 |   32 |         1 |
| [BitcoinAddress](sanitize_benchmark_test.go)          | 10,766,221 |   112.0 |   48 |         1 |
| [BitcoinCashAddress](sanitize_benchmark_test.go)      | 7,910,431  |   151.6 |   48 |         1 |
| [Custom](sanitize_benchmark_test.go) _(Legacy)_       | 920,336    | 1,277.0 |  944 |        17 |
| [CustomCompiled](sanitize_benchmark_test.go)          | 1,638,974  |   730.6 |   96 |         5 |
| [Decimal](sanitize_benchmark_test.go)                 | 18,779,281 |   62.74 |   24 |         1 |
| [Domain](sanitize_benchmark_test.go)                  | 4,988,238  |   243.2 |  176 |         3 |
| [Domain_PreserveCase](sanitize_benchmark_test.go)     | 5,707,197  |   210.4 |  160 |         2 |
| [Domain_RemoveWww](sanitize_benchmark_test.go)        | 4,991,971  |   240.4 |  176 |         3 |
| [Email](sanitize_benchmark_test.go)                   | 8,781,903  |   137.2 |   48 |         2 |
| [Email_PreserveCase](sanitize_benchmark_test.go)      | 13,118,786 |   92.15 |   24 |         1 |
| [FirstToUpper](sanitize_benchmark_test.go)            | 65,587,063 |   17.93 |   16 |         1 |
| [FormalName](sanitize_benchmark_test.go)              | 15,207,229 |   78.84 |   24 |         1 |
| [HTML](sanitize_benchmark_test.go)                    | 2,557,639  |   469.4 |   64 |         3 |
| [IPAddress](sanitize_benchmark_test.go)               | 11,802,175 |   101.4 |   48 |         3 |
| [IPAddress_IPV6](sanitize_benchmark_test.go)          | 2,997,530  |   384.0 |  112 |         3 |
| [Numeric](sanitize_benchmark_test.go)                 | 27,050,888 |    44.0 |   16 |         1 |
| [PhoneNumber](sanitize_benchmark_test.go)             | 9,693,738  |   127.9 |   24 |         1 |
| [PathName](sanitize_benchmark_test.go)                | 15,465,885 |   78.74 |   24 |         1 |
| [Punctuation](sanitize_benchmark_test.go)             | 9,166,885  |   130.7 |   48 |         1 |
| [ScientificNotation](sanitize_benchmark_test.go)      | 19,580,979 |   61.32 |   24 |         1 |
| [Scripts](sanitize_benchmark_test.go)                 | 2,025,324  |   594.6 |   64 |         2 |
| [SingleLine](sanitize_benchmark_test.go)              | 12,599,416 |   95.94 |   32 |         1 |
| [Time](sanitize_benchmark_test.go)                    | 24,114,907 |   48.93 |   16 |         1 |
| [URI](sanitize_benchmark_test.go)                     | 11,414,026 |   104.7 |   32 |         1 |
| [URL](sanitize_benchmark_test.go)                     | 11,462,407 |   105.1 |   32 |         1 |
| [XML](sanitize_benchmark_test.go)                     | 4,179,268  |   288.5 |   56 |         3 |
| [XSS](sanitize_benchmark_test.go)                     | 3,499,938  |   345.1 |   40 |         2 |

> These benchmarks reflect fast, allocation-free lookups for most retrieval functions, ensuring optimal performance in production environments.

<br/>

## 🛠️ Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## 🤖 AI Compliance
This project documents expectations for AI assistants using a few dedicated files:

- [AGENTS.md](.github/AGENTS.md) — canonical rules for coding style, workflows, and pull requests used by [Codex](https://chatgpt.com/codex).
- [CLAUDE.md](.github/CLAUDE.md) — quick checklist for the [Claude](https://www.anthropic.com/product) agent.
- [.cursorrules](.cursorrules) — machine-readable subset of the policies for [Cursor](https://www.cursor.so/) and similar tools.
- [sweep.yaml](.github/sweep.yaml) — rules for [Sweep](https://github.com/sweepai/sweep), a tool for code review and pull request management.

Edit `AGENTS.md` first when adjusting these policies, and keep the other files in sync within the same pull request.

<br/>

## 👥 Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |

<br/>

## 🤝 Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap:
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-sanitize&utm_term=go-sanitize&utm_content=go-sanitize) to ensure this journey continues indefinitely! :rocket:


[![Stars](https://img.shields.io/github/stars/mrz1836/go-sanitize?label=Please%20like%20us&style=social)](https://github.com/mrz1836/go-sanitize/stargazers)

<br/>

## 📝 License

[![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat)](LICENSE)
