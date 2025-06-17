# go-sanitize
> Simple library of sanitation methods for data sanitation and reduction

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
- Review [benchmark results](sanitize_test.go) to assess performance characteristics
- Examine the comprehensive [test suite](sanitize_test.go) for validation and coverage

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
- [`Custom`](sanitize.go): Use a custom regex to filter input
- [`CustomCompiled`](sanitize.go): Use a precompiled regex to filter input
- [`Decimal`](sanitize.go): Keep only decimal or float characters
- [`Domain`](sanitize.go): Sanitize domain, optionally preserving case and removing www
- [`Email`](sanitize.go): Normalize an email address
- [`FirstToUpper`](sanitize.go): Capitalize the first letter of a string
- [`FormalName`](sanitize.go): Keep only formal name characters
- [`HTML`](sanitize.go): Strip HTML tags
- [`IPAddress`](sanitize.go): Return sanitized IPv4 or IPv6 address
- [`Numeric`](sanitize.go): Remove all but numeric digits
- [`PathName`](sanitize.go): Sanitize to a path-friendly name
- [`Punctuation`](sanitize.go): Allow letters, numbers and basic punctuation
- [`ScientificNotation`](sanitize.go): Keep characters valid in scientific notation
- [`Scripts`](sanitize.go): Remove script, iframe and object tags
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

Run the Go [benchmarks](sanitize_test.go):

```bash script
make bench
```

<br/>

Performance benchmarks for the core functions in this library, executed on an Apple M1 Max (ARM64):

### Benchmark Results

| Benchmark                                   | Iterations |   ns/op | B/op | allocs/op |
|---------------------------------------------|------------|--------:|-----:|----------:|
| [Alpha](sanitize_test.go)                   | 1,876,178  |   630.1 |  120 |         6 |
| [Alpha_WithSpaces](sanitize_test.go)        | 2,686,694  |   447.4 |   88 |         4 |
| [AlphaNumeric](sanitize_test.go)            | 1,598,070  |   759.4 |  128 |         6 |
| [AlphaNumeric_WithSpaces](sanitize_test.go) | 1,963,266  |   621.1 |  112 |         4 |
| [BitcoinAddress](sanitize_test.go)          | 2,151,312  |   552.5 |  161 |         4 |
| [BitcoinCashAddress](sanitize_test.go)      | 1,615,339  |   738.0 |  160 |         4 |
| [Custom](sanitize_test.go)                  | 920,336    | 1,277.0 |  944 |        17 |
| [CustomCompiled](sanitize_test.go)          | 1,638,974  |   730.6 |   96 |         5 |
| [Decimal](sanitize_test.go)                 | 2,046,079  |   591.1 |   56 |         3 |
| [Domain](sanitize_test.go)                  | 2,537,883  |   470.8 |  226 |         6 |
| [Domain_PreserveCase](sanitize_test.go)     | 2,880,139  |   420.0 |  209 |         5 |
| [Domain_RemoveWww](sanitize_test.go)        | 1,671,598  |   718.3 |  274 |         9 |
| [Email](sanitize_test.go)                   | 2,159,338  |   555.8 |  137 |         6 |
| [Email_PreserveCase](sanitize_test.go)      | 2,641,934  |   453.1 |  112 |         5 |
| [FirstToUpper](sanitize_test.go)            | 13,212,590 |    90.3 |   24 |         1 |
| [FirstToUpperBuilder](sanitize_test.go)     | 65,667,067 |    18.1 |   16 |         1 |
| [FormalName](sanitize_test.go)              | 3,332,299  |   361.7 |   64 |         3 |
| [HTML](sanitize_test.go)                    | 2,557,639  |   469.4 |   64 |         3 |
| [IPAddress](sanitize_test.go)               | 2,936,395  |   407.8 |   80 |         5 |
| [IPAddress_IPV6](sanitize_test.go)          | 1,000,000  | 1,066.0 |  225 |         6 |
| [Numeric](sanitize_test.go)                 | 2,952,349  |   410.1 |   40 |         3 |
| [PathName](sanitize_test.go)                | 2,336,929  |   510.0 |   64 |         3 |
| [Punctuation](sanitize_test.go)             | 1,895,738  |   621.8 |  160 |         4 |
| [ScientificNotation](sanitize_test.go)      | 1,956,897  |   612.8 |   56 |         3 |
| [Scripts](sanitize_test.go)                 | 2,025,324  |   594.6 |   64 |         2 |
| [SingleLine](sanitize_test.go)              | 555,826    | 2,141.0 |   96 |         4 |
| [Time](sanitize_test.go)                    | 2,183,936  |   549.7 |   40 |         3 |
| [URI](sanitize_test.go)                     | 2,319,432  |   516.5 |   80 |         3 |
| [URL](sanitize_test.go)                     | 2,322,772  |   515.5 |   80 |         3 |
| [XML](sanitize_test.go)                     | 4,179,268  |   288.5 |   56 |         3 |
| [XSS](sanitize_test.go)                     | 3,499,938  |   345.1 |   40 |         2 |

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
