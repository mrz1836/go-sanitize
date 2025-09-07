# AGENTS.md

## 🎯 Purpose & Scope

This file defines the **baseline standards, workflows, and structure** for *all contributors and AI agents* operating within this repository. It serves as the root authority for engineering conduct, coding conventions, and collaborative norms.

It is designed to help AI assistants (e.g., Codex, Claude, Cursor, Sweep AI) and human developers alike understand our practices, contribute clean and idiomatic code, and navigate the codebase confidently and effectively.

> Whether reading, writing, testing, or committing code, **you must adhere to the rules in this document.**

Additional `AGENTS.md` files **may exist in subdirectories** to provide more contextual or specialized guidance. These local agent files are allowed to **extend or override** the root rules to fit the needs of specific packages, services, or engineering domains—while still respecting the spirit of consistency and quality defined here.

<br/>

---

<br/>

## 🔍 Project Overview

`go-sanitize` is a specialized Go library for sanitizing and validating user‑provided data, ensuring that web applications and APIs process secure, normalized, and trusted inputs. It includes functions for:

- Removing or escaping dangerous HTML and script‑related tags
- Validating and formatting email addresses and URLs
- Stripping or trimming unwanted characters (including whitespace and special symbols)
- Escaping sequences that could lead to XSS (Cross‑Site Scripting) or SQL injection
- Normalizing various data types (e.g., time strings, XML content, URIs)

By using `go-sanitize`, developers can greatly reduce security risks from untrusted sources, handle complex input edge cases, and maintain clarity and consistency across Go projects.

<br/>

---

<br/>

## 📁 Directory Structure
| Directory        | Description                                                                       |
|------------------|-----------------------------------------------------------------------------------|
| `.github/`       | Issue templates, workflows, and community documentation                           |
| `.vscode/`       | VS Code settings and extensions for development                                   |
| `.make/`         | Shared Makefile targets used by `Makefile`                                        |
| `examples/`      | Example program demonstrating package usage                                       |
| `.` (root)       | Source files and tests for the local package                                      |

<br/>

---

<br/>

### 📚 Related Governance Documents

For more detailed guidance and supporting documentation, refer to the following project-level resources:

* `CITATION.cff` — Metadata for citing this project; GitHub uses it to render citation information
* `CODEOWNERS` - Ownership of the repository and various directories
* `CODE_OF_CONDUCT.md` — Expected behavior and enforcement
* `CODE_STANDARDS.md` — Style guides and best practices
* `CONTRIBUTING.md` — Guidelines for contributing code, issues, and ideas
* `README.md` — Project overview, goals, and setup instructions
* `SECURITY.md` — Vulnerability reporting and security process
* `SUPPORT.md` — How to get help or request new features

<br/>

---

<br/>

## 🎯 Go Essentials

These are the non-negotiable practices that define professional Go development. Every function, package, and design decision should reflect these principles.

<br/><br/>

### 🌐 Context-First Design

Context should flow through your entire call stack—no exceptions.

* **Always pass `context.Context` as the first parameter** for any operation that could be canceled, timeout, or carry request-scoped values
* **Never store context in structs**—pass it explicitly through function calls
* **Use `context.Background()` only at the top level** (main, tests, or service initialization)
* **Derive child contexts** using `context.WithTimeout()`, `context.WithCancel()`, or `context.WithValue()`
* **Respect context cancellation** by checking `ctx.Done()` in long-running operations

```go
// ✅ Correct: Context as first parameter
func ProcessUserData(ctx context.Context, userID string) error {
    // Check for cancellation before expensive operations
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }

    // Pass context down the call chain
    return database.FetchUser(ctx, userID)
}

// 🚫 Incorrect: No context parameter
func ProcessUserData(userID string) error {
    return database.FetchUser(userID) // Can't be canceled or timeout
}
```

<br/>

---

<br/>

### 🔌 Interface Design Philosophy

Interfaces define contracts, not implementations. Keep them minimal and focused.

* **Accept interfaces, return concrete types** — caller decides what they need; you provide specific value
* **Keep interfaces small** — prefer single-method interfaces (`io.Reader`, `io.Writer`, `io.Closer`)
* **Define interfaces where they're used**, not where they're implemented (consumer-driven)
* **Use composition over large interfaces** — combine small interfaces when needed
* **Name single-method interfaces with `-er` suffix** (`Reader`, `Writer`, `Validator`)

```go
// ✅ Small, focused interface defined at point of use
type UserValidator interface {
    ValidateUser(ctx context.Context, user User) error
}

func ProcessSignup(ctx context.Context, validator UserValidator, user User) error {
    if err := validator.ValidateUser(ctx, user); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    // Process signup...
    return nil
}

// 🚫 Large, monolithic interface
type UserService interface {
    ValidateUser(ctx context.Context, user User) error
    CreateUser(ctx context.Context, user User) error
    UpdateUser(ctx context.Context, user User) error
    DeleteUser(ctx context.Context, userID string) error
    ListUsers(ctx context.Context) ([]User, error)
    // ... 15 more methods
}
```

<br/>

---

<br/>

### ⚡ Goroutine Discipline

Goroutines are cheap to create but expensive to debug when mismanaged.

* **Always have a clear lifecycle** — know when goroutines start and how they terminate
* **Use context for cancellation** — never leave goroutines hanging
* **Avoid naked `go func()`** — wrap in functions that handle errors and cleanup
* **Use `sync.WaitGroup` or channels** for coordination and synchronization
* **Handle goroutine panics** — use `defer recover()` for background workers
* **Limit concurrency** — use worker pools or semaphores to prevent resource exhaustion

```go
// ✅ Well-managed goroutine with proper lifecycle
func ProcessBatch(ctx context.Context, items []Item) error {
    var wg sync.WaitGroup
    errCh := make(chan error, len(items))

    for _, item := range items {
        wg.Add(1)
        go func(item Item) {
            defer wg.Done()
            defer func() {
                if r := recover(); r != nil {
                    errCh <- fmt.Errorf("panic processing item %v: %v", item.ID, r)
                }
            }()

            select {
            case <-ctx.Done():
                errCh <- ctx.Err()
                return
            default:
            }

            if err := processItem(ctx, item); err != nil {
                errCh <- fmt.Errorf("failed to process item %v: %w", item.ID, err)
            }
        }(item)
    }

    wg.Wait()
    close(errCh)

    for err := range errCh {
        if err != nil {
            return err // Return first error encountered
        }
    }
    return nil
}

// 🚫 Unmanaged goroutine
func ProcessBatch(items []Item) {
    for _, item := range items {
        go func(item Item) {
            processItem(item) // No error handling, no cancellation, no lifecycle
        }(item)
    }
    // No way to know when processing is complete
}
```

<br/>

---

<br/>

### 🚫 No Global State

Global variables make code unpredictable, hard to test, and create hidden dependencies.

* **No package-level variables** that hold mutable state
* **Use dependency injection** — pass dependencies explicitly through constructors
* **Prefer `struct` fields over globals** for configuration and state
* **Use `context.Context`** for request-scoped values instead of globals
* **Constants and immutable data are acceptable** at package level

```go
// ✅ Dependency injection pattern
type UserService struct {
    db     Database
    logger Logger
    config Config
}

func NewUserService(db Database, logger Logger, config Config) *UserService {
    return &UserService{
        db:     db,
        logger: logger,
        config: config,
    }
}

func (s *UserService) CreateUser(ctx context.Context, user User) error {
    s.logger.Info("creating user", "userID", user.ID)
    return s.db.Insert(ctx, user)
}

// 🚫 Global state
var (
    globalDB     Database
    globalLogger Logger
    globalConfig Config
)

func CreateUser(ctx context.Context, user User) error {
    globalLogger.Info("creating user", "userID", user.ID) // Hidden dependency
    return globalDB.Insert(ctx, user)                     // Hard to test
}
```

<br/>

---

<br/>

### 🚫 No `init()` Functions

`init()` functions create unpredictable initialization order and hidden side effects.

* **Use explicit constructors** (`NewXxx()` functions) instead of `init()`
* **Make initialization lazy** — initialize on first use when possible
* **Use `sync.Once`** for one-time initialization that must happen exactly once
* **Prefer dependency injection** over package-level initialization
* **Exception**: Only use `init()` for registering with external systems (e.g., database drivers, but even then, prefer explicit registration)

```go
// ✅ Explicit initialization
type Cache struct {
    mu    sync.RWMutex
    data  map[string]interface{}
    once  sync.Once
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

func (c *Cache) ensureInitialized() {
    c.once.Do(func() {
        // One-time setup that's expensive
        c.data = loadInitialData()
    })
}

// 🚫 Hidden initialization
var globalCache map[string]interface{}

func init() {
    globalCache = make(map[string]interface{})
    // This runs at import time - unpredictable order
    // Hard to test, hard to control
}
```

<br/>

---

<br/>

### ⚠️ Error Handling Excellence

* **Always check errors**
* **Use `if err != nil { return err }`** for early returns
* **Use `errors.Is()`** or `errors.As()` for error type checks
* **Use `fmt.Errorf` for wrapping errors** with context
* **Prefer `errors.New()`** over `fmt.Errorf`
* **Use custom error types sparingly**
* **Avoid returning ambiguous errors;** provide context
* **Avoid using `panic`** for expected errors; reserve it for unrecoverable situations
* **Use `errors.Unwrap()`** to access underlying errors when needed
* **Use `errors.Join()`** to combine multiple errors when appropriate
* **Wrap errors with context** using `fmt.Errorf("operation failed: %w", err)`
* **Return early on errors** — avoid deep nesting with guard clauses
* **Log errors at the boundary** — don't log the same error multiple times as it bubbles up

```go
// ✅ Comprehensive error handling
func ProcessPayment(ctx context.Context, payment Payment) error {
    if payment.Amount <= 0 {
        return errors.New("payment amount must be positive")
    }

    user, err := userRepo.GetUser(ctx, payment.UserID)
    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            return fmt.Errorf("cannot process payment for unknown user %s", payment.UserID)
        }
        return fmt.Errorf("failed to fetch user %s: %w", payment.UserID, err)
    }

    if err := validatePaymentMethod(ctx, payment.Method); err != nil {
        return fmt.Errorf("invalid payment method: %w", err)
    }

    txn, err := chargePayment(ctx, payment)
    if err != nil {
        return fmt.Errorf("payment charge failed for user %s: %w", user.ID, err)
    }

    if err := auditRepo.LogTransaction(ctx, txn); err != nil {
        // Log but don't fail the payment
        log.Error("failed to audit transaction", "txnID", txn.ID, "error", err)
    }

    return nil
}

// 🚫 Poor error handling
func ProcessPayment(ctx context.Context, payment Payment) error {
    user, _ := userRepo.GetUser(ctx, payment.UserID) // Ignored error

    validatePaymentMethod(ctx, payment.Method) // Ignored return value

    txn, err := chargePayment(ctx, payment)
    if err != nil {
        return err // No context about what failed
    }

    auditRepo.LogTransaction(ctx, txn) // Ignored error
    return nil
}
```

<br/>

---

<br/>

### 📦 Module Hygiene

Go modules are the foundation of dependency management and reproducible builds.

* **Always use Go modules** — never develop outside a module
* **Pin dependencies to specific versions** — avoid `latest` or floating versions in production
* **Use `go mod tidy`** after any dependency changes to clean up unused dependencies
* **Prefer minimal module graphs** — avoid deep dependency trees when possible
* **Use `replace` directives sparingly** — only for development or emergency patches
* **Document major version upgrades** — breaking changes should be called out in PRs

```bash
# ✅ Proper module management workflow
go mod init github.com/company/project
go get github.com/stretchr/testify@v1.8.4  # Pin to specific version
go mod tidy                                  # Clean up go.mod and go.sum
go mod verify                               # Verify dependencies haven't been tampered with

# ✅ Check for vulnerabilities
govulncheck ./...

# ✅ Update dependencies (with care)
go get -u ./...  # Update to latest minor/patch versions
go mod tidy

# ✅ Use Makefile for common tasks
make update # updates dependencies and runs go mod tidy

# 🚫 Avoid using `replace` unless absolutely necessary
# go.mod
replace github.com/some/dependency => github.com/some/dependency v1.2.3

```

<br/>

---

<br/>

### 🔧 Performance & Profiling

Write code that performs well by default, and measure when optimization is needed.

* **Use `make bench`** to establish performance baselines
* **Profile with `go tool pprof`** when investigating performance issues
* **Avoid premature optimization** — write clear code first, optimize bottlenecks later
* **Use benchmarks** to validate that optimizations actually improve performance
* **Pool expensive objects** with `sync.Pool` when allocation pressure is high
* **Consider memory allocation patterns** — prefer slices to maps for small datasets

```go
// ✅ Performance-conscious code with benchmarks
func BenchmarkUserProcessing(b *testing.B) {
    users := generateTestUsers(1000)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        processUsers(users)
    }
}

func processUsers(users []User) []ProcessedUser {
    // Pre-allocate slice to avoid repeated allocations
    result := make([]ProcessedUser, 0, len(users))

    for _, user := range users {
        processed := ProcessedUser{
            ID:   user.ID,
            Name: strings.ToUpper(user.Name), // Simple transformation
        }
        result = append(result, processed)
    }

    return result
}
```

<br/>

---

<br/>

### 🛠 Formatting & Linting

Code must be cleanly formatted and pass all linters before being committed.

```bash
go fmt ./...
goimports -w .
gofumpt -w ./...
golangci-lint run
go vet ./...
```

> Refer to `.golangci.json` for the full set of enabled linters and formatters.

Editors should honor `.editorconfig` for indentation and whitespace rules, and
Git respects `.gitattributes` to enforce consistent line endings across
platforms.

<br/>

---

<br/>

### 🧪 Testing Standards

We use the `testify` suite for unit tests. All tests must follow these conventions:

* Name tests using the pattern: `TestFunctionNameScenarioDescription` (no underscores) (PascalCase)
* Use `testify` when possible, do not use `testing` directly
* Use `testify/assert` for general assertions
* Use `testify/require` for:
	* All error or nil checks
	* Any test where failure should halt execution
	* Any test where a pointer or complex structure is required to be used after the check
* Use `require.InDelta` or `require.InEpsilon` for floating-point comparisons
* Prefer **table-driven tests** for clarity and reusability, always have a name for each test case
* Use subtests (`t.Run`) to isolate and describe scenarios
* If the test is in a test suite, always use the test suite instead of `t` directly
* **Optionally use** `t.Parallel()` , but try and avoid it unless testing for concurrency issues
* Avoid flaky, timing-sensitive, or non-deterministic tests
* Mock external dependencies — tests should be fast and deterministic
* Use descriptive test names that explain the scenario being tested
* Test error cases — ensure your error handling actually works

Run tests locally with:

```bash
go test ./...
```

> All tests must pass in CI prior to merge.

<br/>

---

<br/>

### 🔍 Fuzz Tests (Optional)

Fuzz tests help uncover unexpected edge cases by generating random inputs. While not required, they are encouraged for **small, self-contained functions**.

Best practices:
* Keep fuzz targets short and deterministic
* Seed the corpus with meaningful values
* Run fuzzers with `go test -fuzz=. -run=^$` when exploring edge cases
* Limit iterations for local runs to maintain speed

<br/>

---

<br/>

## 🛠 Makefile Overview

The repository's `Makefile` includes reusable targets from `.make/common.mk` and
`.make/go.mk`. The root file exposes a few high-level commands while the files
under `.make` contain the bulk of the build logic.

`common.mk` provides utility tasks for releasing with GoReleaser, tagging
releases, and updating the releaser tool. It also offers the `diff` and `help`
commands used across projects.

`go.mk` supplies Go-specific helpers for linting, testing, generating code,
building binaries, and updating dependencies. Targets such as `lint`, `test`,
`test-ci`, and `coverage` are defined here and invoked by the root `Makefile`.

Use `make help` to view the full list of supported commands.

<br/>

---

<br/>

## 📈 Code Coverage

* Code coverage thresholds and rules are defined in `codecov.yml`
* Aim to provide meaningful test coverage for all new logic and edge cases
* Cover every public function with at least one test
* Aim for >= 90% coverage across the codebase (ideally 100%)
* Use `go test -coverprofile=coverage.out ./...` to generate coverage reports

<br/>

---

<br/>

## ✍️ Naming Conventions

Follow Go naming idioms and the standards outlined in [Effective Go](https://go.dev/doc/effective_go):

<br/><br/>

### Packages

* Short, lowercase, one-word (e.g., `auth`, `rpc`, `block`)
* Avoid `util`, `common`, or `shared`
* Exception: standard lib wrapper like `httputil`
* Must have a clear concise package comment in a .go file with the same name as the package

### Files

* Naming using: snake_case (e.g., `block_header.go`, `test_helper.go`)
* Go file names are lowercase
* Test files: `_test.go`
* Generated files: annotate with a `// Code generated by go generate; DO NOT EDIT.` header

### Functions & Methods

* `VerbNoun` naming (e.g., `CalculateHash`, `ReadFile`)
* Constructors: `NewXxx` or `MakeXxx`
* Getters: field name only (`Name()`)
* Setters: `SetXxx(value)`

### Variables

* Exported: `CamelCase` (e.g., `HTTPTimeout`)
* Internal: `camelCase` (e.g., `localTime`)
* Idioms: `i`, `j`, `err`, `tmp` accepted

### Interfaces

* Single-method: use `-er` suffix (e.g., `Reader`, `Closer`)
* Multi-method: use role-based names (e.g., `FileSystem`, `StateManager`)

<br/>

---

<br/>

## 📘 Commenting Standards

Great engineers write great comments. You're not here to state the obvious—you're here to document decisions, highlight edge cases, and make sure the next dev (or AI) doesn't repeat your mistakes.

<br/><br/>

### 🧠 Guiding Principles

* **Comment the "why", not the "what"**

  > The code already tells us *what* it's doing. Your job is to explain *why* it's doing it that way—especially if it's non-obvious, nuanced, or a workaround.

* **Explain side effects, caveats, and constraints**

  > If the function touches global state, writes to disk, mutates shared memory, or makes assumptions—write it down.

* **Don't comment on broken code—fix or delete it**

  > Dead or disabled code with TODOs are bad signals. If it's not worth fixing now, delete it and add an issue instead.

* **Your comments are part of the product**

  > Treat them like UX copy. Make them clear, concise, and professional. You're writing for peers, not compilers.

<br/><br/>

### 🔤 Function Comments (Exported)

Every exported function **must** include a Go-style comment that:

* Starts with the function name
* States its purpose clearly
* Documents:
  * **Steps**: Include if the function performs a non-obvious sequence of operations.
  * **Parameters**: Always describe all parameters when present.
  * **Return values**: Document return types and their meaning if not trivially understood.
  * **Side effects**: Note any I/O, database writes, external calls, or mutations that aren't local to the function.
  * **Notes**: Include any assumptions, constraints, or important context that the caller should know.

Here is a template for function comments that is recommended to use:

```go
// FunctionName does [what this function does] in [brief context].
//
// This function performs the following steps: [if applicable, describe the main steps in a bullet list]
// - [First major action or check performed]
// - [Second action or branching logic explained, if relevant]
//    - [Details about possible outcomes or internal branching]
// - [Additional steps with sub-bullets as needed]
// - [Final steps and cleanup, if applicable]
//
// Parameters:
// - ctx: [Purpose of context in this function]
// - paramName: [Explanation of each parameter and what it controls or affects]
//
// Returns:
// - [What is returned; error behavior if any]
//
// Side Effects:
// - [Any side effects, such as modifying global state, writing to disk, etc.]
//
// Notes:
// - [Caveats, assumptions, or recommendations—e.g., transaction usage, concurrency, etc.]
// - [Any implicit contracts with the caller or system constraints]
// - [Mention if this function is part of a larger workflow or job system]
```

<br/><br/>

### 📦 Package-Level Comments

* Each package **must** include a package-level comment in a file named after the package (e.g., `auth.go` for package `auth`).
* If no logical file fits, add a `doc.go` with the comment block.
* Use it to explain:
    * The package purpose
    * High-level API boundaries
    * Expected use-cases and design notes

Here is a template for package comments that is recommended to use:

```go
// Package PackageName provides [brief description of the package's purpose].
//
// This package implements [high-level functionality or API] and is designed to [describe intended use cases].
//
// Key features include: [if applicable, list key features or components]
// - [Feature 1: brief description of what it does]
// - [Feature 2: brief description of what it does]
// - [Feature 3: brief description of what it does]
//
// The package is structured to [explain any architectural decisions, e.g., modularity, separation of concerns].
// It relies on [mention any key dependencies or external systems].
//
// Usage examples:
// [Provide a simple example of how to use the package, if applicable]
//
// Important notes:
// - [Any important assumptions or constraints, e.g., concurrency model, state management]
// - [Any known limitations or edge cases]
// - [Any specific configuration or initialization steps required]
//
// This package is part of the larger [Project Name] ecosystem and interacts with [mention related packages or systems].
package PackageName
```

<br/><br/>

### 🧱 Inline Comments

Use inline comments **strategically**, not excessively.

* Use them to explain "weird" logic, workarounds, or business rules.
* Prefer **block comments (`//`)** on their own line over trailing comments.
* Avoid obvious noise:

🚫 `i++ // increment i`

✅ `// Skip empty rows to avoid panic on CSV parse`

<br/><br/>

### ⚙️ Comment Style

* Use **complete sentences** with punctuation.
* Keep your tone **precise, confident, and modern**—you're not writing a novel, but you're also not writing legacy COBOL.
* Avoid filler like "simple function" or "just does X".
* Don't leave TODOs unless:
    * They are immediately actionable
    * (or) they reference an issue
    * They include a timestamp or owner

<br/><br/>

### 🧬 AI Agent Directives

If you're an AI contributing code:

* Treat your comments like commit messages—**use active voice, be declarative**
* Use comments to **make intent explicit**, especially for generated or AI-authored sections
* Avoid hallucinating context—if you're unsure, omit or tag with `// AI: review this logic`
* Flag areas of uncertainty or external dependency (e.g., "// AI: relies on external config structure")

<br/><br/>

### 🔥 Comment Hygiene

* Remove outdated comments aggressively.
* Keep comments synced with refactoring.
* Use `//nolint:<linter> // message` only with clear, justified context and explanation.

<br/>

---

<br/>

## 📝 Modifying Markdown Documents

Markdown files (e.g., `README.md`, `AGENTS.md`, `CONTRIBUTING.md`) are first-class citizens in this repository. Edits must follow these best practices:

* **Write with intent** — Be concise, clear, and audience-aware. Each section should serve a purpose.
* **Use proper structure** — Maintain consistent heading levels, spacing, and bullet/number list formatting.
* **Full Table Borders** — Use full borders for tables to ensure readability across different renderers.
* **Table Border Spacing** — Make sure tables have appropriate spacing for clarity.
* **Preserve voice and tone** — Match the professional tone and style used across the project documentation.
* **Preview before committing** — Always verify rendered output locally or in a PR to avoid broken formatting.
* **Update references** — When renaming files or sections, update internal links and the table of contents if present.

> Markdown updates should be treated with the same care as code—clean, purposeful, and reviewed.

<br/>

---

<br/>

## 🔀 Commit & Branch Naming Conventions

Clear history ⇒ easy maintenance. Follow these rules for every commit and branch.

<br/><br/>

### 📌 Commit Message Format

```
<type>(<scope>): <imperative short description>

<body>  # optional, wrap at 72 chars
```

* **`<type>`** — `feat`, `fix`, `docs`, `test`, `refactor`, `chore`, `build`, `ci`
* **`<scope>`** — Affected subsystem or package (e.g., `api`, `deps`). Omit if global.
* **Short description** — ≤ 50 chars, imperative mood ("add pagination", "fix panic")
* **Body** (optional) — What & why, links to issues (`Closes #123`), and breaking‑change note (`BREAKING CHANGE:`)

**Examples**

```
feat(package): add new method called: Thing()
fix(generator): handle malformed JSON input gracefully
docs(README): improve installation instructions
```

> Commits that only tweak whitespace, comments, or docs inside a PR may be squashed; otherwise preserve granular commits.

<br/><br/>

### 📝 Pre-Commit Hooks (Optional)
To ensure consistent commit messages, we use a pre-commit hook that checks the format before allowing a commit. The hook is defined in `.pre-commit-config.yaml` and can be installed with:

```bash
pre-commit install
```

If you don't have `pre-commit` installed, you can install it via Homebrew:
```bash
brew install pre-commit
```

Run the pre-commit hook manually with:
```bash
pre-commit run --all-files
```

> The pre-commit hook will automatically check your commit messages against the defined format and prevent commits that do not comply.

<br/><br/>

### 🌱 Branch Naming

| Purpose            | Prefix      | Example                            |
|--------------------|-------------|------------------------------------|
| Bug Fix            | `fix/`      | `fix/code-off-by-one`              |
| Chore / Meta       | `chore/`    | `chore/upgrade-go-1.24`            |
| Documentation      | `docs/`     | `docs/agents-commenting-standards` |
| Feature            | `feat/`     | `feat/pagination-api`              |
| Hotfix (prod)      | `hotfix/`   | `hotfix/rollback-broken-deploy`    |
| Prototype / Spike  | `proto/`    | `proto/iso3166-expansion`          |
| Refactor / Cleanup | `refactor/` | `refactor/remove-dead-code`        |
| Tests              | `test/`     | `test/generator-edge-cases`        |

* Use **kebab‑case** after the prefix.
* Keep branch names concise yet descriptive.
* PR titles should mirror the branch's purpose (see [✅ Pull Request Conventions](#-pull-request-conventions)).

> CI rely on these prefixes for auto labeling and workflow routing—stick to them.

<br/>

---

<br/>

## ✅ Pull Request Conventions

Pull Requests—whether authored by humans or AI agents—must follow a consistent structure to ensure clarity, accountability, and ease of review.

<br/><br/>

### 🔖 Title Format

```
[Subsystem] Imperative and concise summary of change
```

Examples:

* `[API] Add pagination to client search endpoint`
* `[DB] Migrate legacy rate table schema`
* `[CI] Remove deprecated GitHub Action for testing`

> Use the imperative mood ("Add", "Fix", "Update") to match the style of commit messages and changelogs.

<br/><br/>

### 📝 Pull Request Description

Every PR must include the following **four** sections in the description:

#### 1. **What Changed**

> A clear, bullet‑pointed or paragraph‑level summary of the technical changes.

#### 2. **Why It Was Necessary**

> Context or motivation behind the change. Reference related issues, discussions, or bugs if applicable.

#### 3. **Testing Performed**

> Document:
>
> * Test suites run (e.g., `TestCreateOriginationAccount`)
> * Edge cases covered
> * Manual steps that were taken (if any)

#### 4. **Impact / Risk**

> Call out:
>
> * Breaking changes
> * Regression risk
> * Performance implications
> * Changes in developer experience (e.g., local dev setup, CI time)

<br/><br/>

### 💡 Additional PR Guidelines

* Link related issues with keywords like `Closes #123` or `Fixes #456` if there is a known issue.
* Keep PRs focused and minimal. Prefer multiple small PRs over large ones when possible.
* Use draft PRs early for feedback on in progress work.
* Releases are deployed using **goreleaser**.
* Rules for the release build are located in `.goreleaser.yml` and executed via `.github/workflows/release.yml`.

<br/>

---

<br/>

## 🚀 Release Workflow & Versioning

We follow **Semantic Versioning (✧ SemVer)**:
`MAJOR.MINOR.PATCH` → `1.2.3`

| Segment   | Bumps When …                          | Examples        |
|-----------|---------------------------------------|-----------------|
| **MAJOR** | Breaking API change                   | `1.0.0 → 2.0.0` |
| **MINOR** | Back‑compatible feature / enhancement | `1.2.0 → 1.3.0` |
| **PATCH** | Back‑compatible bug fix / docs        | `1.2.3 → 1.2.4` |

<br/><br/>

### 📦 Tooling

* Releases are driven by **[goreleaser]** and configured in `.goreleaser.yml`.
* Install locally with Homebrew (Mac):
```bash
  brew install goreleaser
````

<br/><br/>

### 🔄 Workflow

| Step | Command                         | Purpose                                                                                            |
|------|---------------------------------|----------------------------------------------------------------------------------------------------|
| 1    | `make release-snap`             | Build & upload a **snapshot** (pre‑release) for quick CI validation.                               |
| 2    | `make tag version=X.Y.Z`        | Create and push a signed Git tag. Triggers GitHub Actions to package the release                   |
| 3    | GitHub Actions                  | CI runs `goreleaser release` on the tag; artifacts and changelog are published to GitHub Releases. |

> **Note for AI Agents:** Do not create or push tags automatically. Only the repository [codeowners](CODEOWNERS) are authorized to tag and publish official releases.

[goreleaser]: https://github.com/goreleaser/goreleaser

<br/>

---

<br/>

## 🏷️ Labeling Conventions (GitHub)

Labels serve as shared vocabulary for categorizing issues, pull requests, and discussions. Proper labeling improves triage, prioritization, automation, and clarity across the engineering lifecycle.

Current labels are located in `.github/labels.yml` and automatically synced into GitHub upon updating the `master` branch.

<br/><br/>

### 🎨 Standard Labels & Usage

| Label Name               | Color     | Description                                                         | When to Use                                                                 |
|--------------------------|-----------|---------------------------------------------------------------------|-----------------------------------------------------------------------------|
| `automerge`              | `#fef2c0` | Safe to merge automatically (e.g., from CI or bot)                  | Label added by automation or trusted reviewers                              |
| `bug-P1`                 | `#b23128` | **Critical bug**, highest priority, impacts all users               | Regressions, major system outages, critical service bugs                    |
| `bug-P2`                 | `#de3d32` | **Moderate bug**, medium priority, affects a subset                 | Broken functionality with known workaround or scoped impact                 |
| `bug-P3`                 | `#f44336` | **Minor bug**, lowest priority, limited user impact                 | Edge case issues, cosmetic UI glitches, legacy bugs                         |
| `chore`                  | `#006b75` | Low-impact, internal tasks                                          | Dependency bumps, code formatting, comment fixes                            |
| `dependabot`             | `#1a70c7` | PR or issue created by or related to Dependabot                     | Automated PRs/issues from Dependabot or related dependency bots             |
| `dependencies`           | `#006b75` | Dependency updates, version bumps, etc.                             | Any PR/issue related to dependency management or upgrades                   |
| `devcontainer`           | `#006b75` | Used for referencing DevContainers                                  | Changes to `.devcontainer` or container-based dev environments              |
| `docker`                 | `#006b75` | Used for referencing Docker related issues                          | Dockerfile, docker-compose, containerization changes                        |
| `documentation`          | `#0075ca` | Improvements or additions to project docs                           | Updates to `README`, onboarding docs, usage guides, code comments           |
| `feature`                | `#0e8a16` | Any new **major feature or capability**                             | Adding new API, CLI command, UI section, or module                          |
| `github-actions`         | `#006b75` | Used for referencing GitHub Actions                                 | Workflow, action, or CI/CD pipeline changes                                 |
| `gomod`                  | `#006b75` | Used for referencing Go Modules                                     | Changes to `go.mod`, `go.sum`, or module management                         |
| `hot-fix`                | `#b60205` | Time-sensitive or production-impacting fix                          | Used with `bug-P1` or urgent code/config changes that must ship immediately |
| `idea`                   | `#cccccc` | Suggestions or brainstorming candidates                             | Feature ideas, process improvements, early-stage thoughts                   |
| `performance`            | `#8bc34a` | Performance improvements or optimizations                           | Optimizing code, reducing latency, improving resource usage                 |
| `prototype`              | `#d4c5f9` | Experimental work that may be unstable or incomplete                | Spike branches, POCs, proof-of-concept work                                 |
| `question`               | `#cc317c` | A request for clarification or feedback                             | Use for technical questions, code understanding, process queries            |
| `refactor`               | `#ffa500` | Non-functional changes to improve structure or readability          | Code cleanups, abstraction, splitting monoliths                             |
| `requires-manual-review` | `#ffd700` | PR or issue requires manual review by a maintainer or security team | Use when automation cannot determine safety or correctness                  |
| `security`               | `#d73a4a` | Security-related issue, vulnerability, or fix                       | Vulnerabilities, security patches, or sensitive changes                     |
| `size/L`                 | `#01579b` | Large change (201–500 lines)                                        | Large PRs, major refactors                                                  |
| `size/M`                 | `#0288d1` | Medium change (51–200 lines)                                        | Moderate PRs, new features, refactors                                       |
| `size/S`                 | `#4fc3f7` | Small change (11–50 lines)                                          | Small, focused PRs                                                          |
| `size/XL`                | `#002f6c` | Very large change (>500 lines)                                      | Massive PRs, sweeping changes                                               |
| `size/XS`                | `#b3e5fc` | Very small change (≤10 lines)                                       | Tiny PRs, typo fixes, single-line changes                                   |
| `stale`                  | `#c2e0c6` | Inactive, obsolete, or no longer relevant                           | Used for automated cleanup or manual archiving of old PRs/issues            |
| `test`                   | `#c2e0c6` | Changes to tests or test infrastructure                             | Unit tests, mocks, integration tests, CI coverage enhancements              |
| `ui-ux`                  | `#fbca04` | Frontend or user experience-related changes                         | CSS/HTML/JS updates, UI behavior tweaks, design consistency                 |
| `update`                 | `#006b75` | General updates not tied to a specific bug or feature               | Routine code changes, small improvements, silent enhancements               |
| `work-in-progress`       | `#fbca04` | Not ready to merge, actively under development                      | Blocks `automerge`, signals in-progress discussion or implementation        |

<br/><br/>

### 🧠 Labeling Best Practices

* Apply labels at the time of PR/issue creation, or during triage.
* Use **only one priority label** (`bug-P1`, `P2`, `P3`) per item.
* Combine labels as needed (e.g., `feature` + `ui-ux` + `test`).
* Don't forget to remove outdated labels (e.g., `work-in-progress` → after merge readiness).
* Use `requires-manual-review` for PRs that need human eyes before merging, especially if they involve security or critical changes.
* Use `automerge` only when the PR is fully approved, tested, and ready for production.
* Use `stale` for issues or PRs that have been inactive for 30+ days, to signal they may need attention or closure.
* Use `size/*` labels to indicate the complexity of the change, which helps reviewers understand the effort involved.
* Use `security` for any PR that addresses vulnerabilities, security patches, or sensitive changes that require extra scrutiny.

<br/>

---

<br/>

## 🧩 CI & Validation

CI automatically runs on every PR to verify:

* Formatting (`go fmt` and `goimports` and `gofumpt`)
* Linting (`make lint`)
* Tests (`make test`)
* Fuzz tests (if applicable) (`make run-fuzz-tests`)
* This codebase uses GitHub Actions; test workflows reside in `.github/workflows/fortress.yml` and `.github/workflows/fortress-test-suite.yml`.
* Pin each external GitHub Action to a **full commit SHA** (e.g., `actions/checkout@2f3b4a2e0e471e13e2ea2bc2a350e888c9cf9b75`) as recommended by GitHub's [security hardening guidance](https://docs.github.com/en/actions/security-guides/security-hardening-for-github-actions#using-pinned-actions). Dependabot will track and update these pinned versions automatically.

Failing PRs will be blocked. AI agents should iterate until CI passes.

<br/>

---

<br/>

## 🔐 Dependency Management

Dependency hygiene is critical for security, reproducibility, and developer experience. Follow these practices to ensure our module stays stable, up to date, and secure.

<br/><br/>

### 📦 Module Management

* All dependencies must be managed via **Go Modules** (`go.mod`, `go.sum`)

* After adding, updating, or removing imports, run:

  ```bash
  go mod tidy
  ```

* Periodically refresh dependencies with:

  ```bash
  go get -u ./...
  ```

> Avoid unnecessary upgrades near release windows—review major version bumps carefully for breaking changes.

<br/><br/>

### 🛡️ Security Scanning

* Use [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) to identify known vulnerabilities:
```bash
  govulncheck ./...
```

* Run via make command:
```bash
  make govulncheck
```

* Run [gitleaks](https://github.com/gitleaks/gitleaks) before committing code to detect hardcoded secrets or sensitive data in the repository:
```bash
brew install gitleaks
gitleaks detect --source . --log-opts="--all" --verbose
```

* Run [OSSAR](https://github.com/github/ossar-action) to perform additional static analysis security checks (see `.github/workflows/ossar.yml`).

* Address critical advisories before merging changes into `master`

* Document any intentionally ignored vulnerabilities with clear justification and issue tracking

* We follow the [OpenSSF](https://openssf.org) best practices to ensure this repository remains compliant with industry‑standard open source security guidelines

<br/><br/>

### 📁 Version Control

* Never manually edit `go.sum`
* Do not vendor dependencies; we rely on modules for reproducibility
* Lockstep upgrades across repos (when applicable) should be coordinated and noted in PRs

> Changes to dependencies should be explained in the PR description and ideally linked to the reason (e.g., bug fix, security advisory, feature requirement).

<br/>

---

<br/>

## 🛡️Security Considerations & Vulnerability Reporting

Security is a first-class requirement. If you discover a vulnerability—no matter how small—follow our responsible disclosure process:

* **Do not** open a public issue or pull request.
* Follow the instructions in [`SECURITY.md`](SECURITY.md).
* Include:
  * A clear, reproducible description of the issue
  * Proof‑of‑concept code or steps (if possible)
  * Any known mitigations or workarounds
* You will receive an acknowledgment within **72 hours** and status updates until the issue is resolved.

> For general hardening guidance (e.g., `govulncheck`, dependency pinning), see the [🔐Dependency Management](#-dependency-management) section.

<br/>

---

<br/>

## ⚙️ GitHub Workflows Development

GitHub Actions workflows are critical infrastructure that automate our CI/CD pipeline, testing, security scanning, and release processes. All workflows must be reliable, secure, and maintainable to ensure consistent developer experience and protect our codebase.

<br/><br/>

### 🎯 Workflow Development Priorities

* **Security First** — Use minimal permissions, pin all actions to full commit SHA, and never expose secrets unnecessarily
* **Reliability** — Workflows should be deterministic, idempotent, and handle failure gracefully
* **Performance** — Optimize for speed with appropriate caching, parallel execution, and efficient resource usage
* **Maintainability** — Write self-documenting workflows with clear structure, consistent naming, and comprehensive comments
* **Native GitHub Features** — Prefer built-in GitHub Actions and features over third-party alternatives when possible
* **Least Privilege** — Grant only the minimum permissions required for each job and step
* **Fail Fast** — Design workflows to catch issues early and provide clear, actionable error messages
* **Documentation** — Every workflow must include purpose, triggers, and maintainer information in header comments

<br/><br/>

### 📋 Workflow Development Guidelines

* **Action Pinning** — Pin all external actions to full commit SHA (e.g., `actions/checkout@2f3b4a2e0e471e13e2ea2bc2a350e888c9cf9b75`) for security and reproducibility
* **Permissions** — Use `permissions: read-all` as default, then grant specific write permissions only where needed
* **Concurrency Control** — Include concurrency groups to prevent resource conflicts and optimize runner usage
* **Environment Variables** — Use `env` blocks at appropriate levels (workflow, job, or step) to maintain clarity
* **Error Handling** — Use `continue-on-error` and `if` conditionals strategically to handle expected failures
* **Matrix Builds** — Leverage matrix strategies for testing across multiple versions, platforms, or configurations
* **Caching** — Implement intelligent caching for dependencies, build artifacts, and test results
* **Secrets Management** — Use GitHub Secrets for sensitive data; never hardcode credentials or tokens
* **Branch Protection** — Ensure critical workflows are required status checks before merging
* **Workflow Naming** — Use kebab-case names that clearly describe the workflow purpose

<br/><br/>

### 🏗️ Workflow Template

Use this template as the foundation for all new GitHub Actions workflows:

```yaml
# ------------------------------------------------------------------------------
#  [Workflow Name] Workflow
#
#  Purpose: [Brief description of what this workflow does and why it exists]
#
#  Triggers: [List the events that trigger this workflow]
#
#  Maintainer: @[github-username]
# ------------------------------------------------------------------------------

name: [workflow-name]

# ————————————————————————————————————————————————————————————————
# Trigger Configuration
# ————————————————————————————————————————————————————————————————
on:
  [trigger-events]

# ————————————————————————————————————————————————————————————————
# Permissions
# ————————————————————————————————————————————————————————————————
permissions: read-all

# ————————————————————————————————————————————————————————————————
# Concurrency Control
# ————————————————————————————————————————————————————————————————
concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  [job-name]:
    runs-on: ubuntu-latest
    permissions:
      [specific-permissions]: [read|write]
    steps:
      # ————————————————————————————————————————————————————————————————
      # [Step Category/Purpose]
      # ————————————————————————————————————————————————————————————————
      - name: [Step Name]
        uses: [action@full-commit-sha]
        with:
          [parameters]

      # ————————————————————————————————————————————————————————————————
      # [Next Step Category/Purpose]
      # ————————————————————————————————————————————————————————————————
      - name: [Next Step Name]
        run: |
          [commands]
        env:
          [environment-variables]
```

<br/><br/>

### 🔍 Template Usage Notes

* Replace all bracketed placeholders `[...]` with actual values
* Use descriptive names for workflows, jobs, and steps
* Include section separators (dashes) to organize logical groups of steps
* Add environment variables in the appropriate scope (workflow, job, or step level)
* Document complex logic with inline comments using `#` within run blocks
* Test workflows thoroughly in draft PRs before merging to avoid CI/CD disruption

<br/><br/>

### 🚦 Workflow Validation Checklist

Before merging any workflow changes, verify:

* [ ] All external actions are pinned to full commit SHA
* [ ] Permissions follow least-privilege principle
* [ ] Concurrency control prevents resource conflicts
* [ ] Workflow header includes purpose, triggers, and maintainer
* [ ] Section separators organize steps logically
* [ ] Environment variables are properly scoped
* [ ] Error handling covers expected failure scenarios
* [ ] Workflow has been tested in a draft PR
* [ ] Documentation reflects any new workflow dependencies or requirements

<br/>

---

<br/>

## 🕓 Change Log (AGENTS.md)

This section tracks notable updates to `AGENTS.md`, including the date, author, and purpose of each revision.
All contributors are expected to append entries here when making meaningful changes to agent behavior, conventions, or policies.


| Date       | Author   | Summary of Changes                                    |
|------------|----------|-------------------------------------------------------|
| 2025-07-10 | @mrz1836 | Go essentials section, labels, and more!              |
| 2025-07-07 | @mrz1836 | Minor mods in formatting, linting, gitleaks           |
| 2025-06-30 | @mrz1836 | Added pre-commit hook guidelines and config reference |
| 2025-06-27 | @mrz1836 | Adapted to fix this project                           |
> For minor edits (typos, formatting), this log update is optional. For all behavioral or structural changes, log entries are **required**.
