## Default to the repo name if empty
ifndef BINARY_NAME
	override BINARY_NAME=app
endif

## Define the binary name
ifdef CUSTOM_BINARY_NAME
	override BINARY_NAME=$(CUSTOM_BINARY_NAME)
endif

## Set the binary release names
DARWIN=$(BINARY_NAME)-darwin
LINUX=$(BINARY_NAME)-linux
WINDOWS=$(BINARY_NAME)-windows.exe

.PHONY: test lint install

bench:  ## Run all benchmarks in the Go application
	@go test -bench=. -benchmem

build-go:  ## Build the Go application (locally)
	@go build -o bin/$(BINARY_NAME)

clean-mods: ## Remove all the Go mod cache
	@go clean -modcache

coverage: ## Shows the test coverage
	@go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

godocs: ## Sync the latest tag with GoDocs
	@test $(GIT_DOMAIN)
	@test $(REPO_OWNER)
	@test $(REPO_NAME)
	@test $(VERSION_SHORT)
	@curl https://proxy.golang.org/$(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)/@v/$(VERSION_SHORT).info

install: ## Install the application
	@go build -o $$GOPATH/bin/$(BINARY_NAME)

install-go: ## Install the application (Using Native Go)
	@go install $(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)

lint: ## Run the Go lint application
	@if [ "$(shell command -v golint)" = "" ]; then go get -u golang.org/x/lint/golint; fi
	@golint

test: ## Runs vet, lint and ALL tests
	@$(MAKE) vet
	@$(MAKE) lint
	@go test ./... -v

test-short: ## Runs vet, lint and tests (excludes integration tests)
	@$(MAKE) vet
	@$(MAKE) lint
	@go test ./... -v -test.short

test-travis: ## Runs all tests via Travis (also exports coverage)
	@$(MAKE) vet
	@$(MAKE) lint
	@go test ./... -race -coverprofile=coverage.txt -covermode=atomic

test-travis-short: ## Runs unit tests via Travis (also exports coverage)
	@$(MAKE) vet
	@$(MAKE) lint
	@go test ./... -test.short -race -coverprofile=coverage.txt -covermode=atomic

uninstall: ## Uninstall the application (and remove files)
	@test $(BINARY_NAME)
	@test $(GIT_DOMAIN)
	@test $(REPO_OWNER)
	@test $(REPO_NAME)
	@go clean -i $(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)
	@rm -rf $$GOPATH/src/$(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)
	@rm -rf $$GOPATH/bin/$(BINARY_NAME)

update:  ## Update all project dependencies
	@go get -u ./... && go mod tidy

vet: ## Run the Go vet application
	@go vet -v ./...