# go.mk — Makefile for Go projects

# Default binary name
ifndef BINARY_NAME
BINARY_NAME := app
endif

ifdef CUSTOM_BINARY_NAME
BINARY_NAME := $(CUSTOM_BINARY_NAME)
endif

# Platform-specific binaries
DARWIN := $(BINARY_NAME)-darwin
LINUX := $(BINARY_NAME)-linux
WINDOWS := $(BINARY_NAME)-windows.exe

# Go build tags
TAGS :=
ifdef GO_BUILD_TAGS
TAGS := -tags $(GO_BUILD_TAGS)
endif

.PHONY: bench
bench: ## Run all benchmarks in the Go application
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem $(TAGS)

.PHONY: build-go
build-go: ## Build the Go application (locally)
	@echo "Building Go app..."
	@go build -o bin/$(BINARY_NAME) $(TAGS)

.PHONY: clean-mods
clean-mods: ## Remove all the Go mod cache
	@echo "Cleaning Go mod cache..."
	@go clean -modcache

.PHONY: coverage
coverage: ## Show test coverage
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./... $(TAGS) && go tool cover -func=coverage.out

.PHONY: generate
generate: ## Run go generate in the base of the repo
	@echo "Running go generate..."
	@go generate -v $(TAGS)

.PHONY: godocs
godocs: ## Trigger GoDocs tag sync
	@echo "Syndicating to GoDocs..."
	@if [ -z "$(GIT_DOMAIN)" ] || [ -z "$(REPO_OWNER)" ] || [ -z "$(REPO_NAME)" ] || [ -z "$(VERSION_SHORT)" ]; then \
		echo "Missing variables for GoDocs push" && exit 1; \
	fi
	@curl -sSf https://proxy.golang.org/$(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)/@v/$(VERSION_SHORT).info

.PHONY: install
install: ## Install the application binary
	@echo "Installing binary..."
	@go build -o $$GOPATH/bin/$(BINARY_NAME) $(TAGS)

.PHONY: install-go
install-go: ## Install using go install with specific version
	@echo "Installing with go install..."
	@go install $(TAGS) $(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)@$(VERSION_SHORT)

.PHONY: lint
lint: ## Run the golangci-lint application (install if not found)
	@if [ "$(shell which golangci-lint)" = "" ]; then \
		if [ "$(shell command -v brew)" != "" ]; then \
			echo "Brew detected, attempting to install golangci-lint..."; \
			if ! brew list golangci-lint &>/dev/null; then \
				brew install golangci-lint; \
			else \
				echo "golangci-lint is already installed via brew."; \
			fi; \
		else \
			echo "Installing golangci-lint via curl..."; \
			GOPATH=$$(go env GOPATH); \
			if [ -z "$$GOPATH" ]; then GOPATH=$$HOME/go; fi; \
			echo "Installation path: $$GOPATH/bin"; \
			curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin v2.1.6; \
		fi; \
	fi; \
	if [ "$(TRAVIS)" != "" ]; then \
		echo "Travis CI environment detected."; \
	elif [ "$(CODEBUILD_BUILD_ID)" != "" ]; then \
		echo "AWS CodePipeline environment detected."; \
	elif [ "$(GITHUB_WORKFLOW)" != "" ]; then \
		echo "GitHub Actions environment detected."; \
	fi; \
	echo "Running golangci-lint..."; \
	golangci-lint run --verbose

.PHONY: run-fuzz-tests
run-fuzz-tests: ## Run fuzz tests for all packages
	@for pkg in $(shell go list ./...); do \
		for fuzz in $$(go test -list ^Fuzz -run=^$$pkg | grep ^Fuzz); do \
			echo "Running fuzz test $$fuzz in $$pkg"; \
			go test -fuzz=$$fuzz -fuzztime=5s $$pkg || exit 1; \
		done; \
	done

.PHONY: test
test: ## Run lint and all tests
	@$(MAKE) lint
	@echo "Running all tests..."
	@go test ./... -v $(TAGS)
	@$(MAKE) run-fuzz-tests

.PHONY: test-unit
test-unit: ## Runs tests and outputs coverage
	@echo "Running unit tests..."
	@go test ./... -race -coverprofile=coverage.txt -covermode=atomic $(TAGS)

.PHONY: test-short
test-short: ## Run tests excluding integration
	@$(MAKE) lint
	@echo "Running short tests..."
	@go test ./... -v -test.short $(TAGS)

.PHONY: test-ci
test-ci: ## CI full test suite with coverage
	@$(MAKE) lint
	@echo "Running CI tests..."
	@go test ./... -race -coverprofile=coverage.txt -covermode=atomic $(TAGS)
	@$(MAKE) run-fuzz-tests

.PHONY: test-ci-no-race
test-ci-no-race: ## CI test suite without race detector
	@$(MAKE) lint
	@echo "Running CI tests (no race)..."
	@go test ./... -coverprofile=coverage.txt -covermode=atomic $(TAGS)
	@$(MAKE) run-fuzz-tests

.PHONY: test-ci-short
test-ci-short: ## CI unit-only short tests
	@$(MAKE) lint
	@echo "Running CI short tests..."
	@go test ./... -test.short -race -coverprofile=coverage.txt -covermode=atomic $(TAGS)

.PHONY: test-no-lint
test-no-lint: ## Run only tests (no lint)
	@echo "Running tests..."
	@go test ./... -v $(TAGS)

.PHONY: uninstall
uninstall: ## Uninstall the Go binary
	@echo "Uninstalling binary..."
	@test -n "$(BINARY_NAME)"
	@test -n "$(GIT_DOMAIN)"
	@test -n "$(REPO_OWNER)"
	@test -n "$(REPO_NAME)"
	@go clean -i $(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)
	@rm -rf $$GOPATH/bin/$(BINARY_NAME)

.PHONY: update
update: ## Update dependencies
	@echo "Updating dependencies..."
	@go get -u ./... && go mod tidy

.PHONY: update-linter
update-linter: ## Upgrade golangci-lint (macOS only)
	@echo "Upgrading golangci-lint..."
	@brew upgrade golangci-lint

.PHONY: vet
vet: ## Run go vet
	@echo "Running go vet..."
	@go vet -v ./... $(TAGS)

.PHONY: govulncheck-install
govulncheck-install: ## Install govulncheck
	@echo "Installing govulncheck..."
	@go install golang.org/x/vuln/cmd/govulncheck@latest

.PHONY: govulncheck
govulncheck: govulncheck-install ## Scan for vulnerabilities
	@echo "Running govulncheck..."
	@govulncheck ./...