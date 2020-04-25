COVER=go tool cover

## Default Repo Domain
GIT_DOMAIN=github.com

## Set the Github Token
#GITHUB_TOKEN=<your_token>

## Automatically detect the repo owner and repo name
REPO_NAME=$(shell basename `git rev-parse --show-toplevel`)
REPO_OWNER=$(shell git config --get remote.origin.url | sed 's/git@$(GIT_DOMAIN)://g' | sed 's/\/$(REPO_NAME).git//g')

## Set the version(s) (injected into binary)
VERSION=$(shell git describe --tags --always --long --dirty)
VERSION_SHORT=$(shell git describe --tags --always --abbrev=0)

.PHONY: test lint clean release

all: lint test vet

bench:  ## Run all benchmarks in the Go application
	go test -bench ./... -benchmem -v

clean: ## Remove previous builds and any test cache data
	go clean -cache -testcache -i -r
	if [ -d ${DISTRIBUTIONS_DIR} ]; then rm -r ${DISTRIBUTIONS_DIR}; fi

clean-mods: ## Remove all the Go mod cache
	go clean -modcache

coverage: ## Shows the test coverage
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

godocs: ## Sync the latest tag with GoDocs
	curl https://proxy.golang.org/$(GIT_DOMAIN)/$(REPO_OWNER)/$(REPO_NAME)/@v/$(VERSION_SHORT).info

help: ## Show all make commands available
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## Run the Go lint application
	golint

release: ## Full production release (creates release in Github)
	 goreleaser --rm-dist
	 make godocs

release-test: ## Full production test release (everything except deploy)
	 goreleaser --skip-publish --rm-dist

release-snap: ## Test the full release (build binaries)
	goreleaser --snapshot --skip-publish --rm-dist

tag: ## Generate a new tag and push (IE: make tag version=0.0.0)
	test $(version)
	git tag -a v$(version) -m "Pending full release..."
	git push origin v$(version)
	git fetch --tags -f

tag-remove: ## Remove a tag if found (IE: make tag-remove version=0.0.0)
	test $(version)
	git tag -d v$(version)
	git push --delete origin v$(version)
	git fetch --tags

tag-update: ## Update an existing tag to current commit (IE: make tag-update version=0.0.0)
	test $(version)
	git push --force origin HEAD:refs/tags/v$(version)
	git fetch --tags -f

test: ## Runs vet, lint and ALL tests
	make vet
	make lint
	go test ./... -v

test-short: ## Runs vet, lint and tests (excludes integration tests)
	make vet
	make lint
	go test ./... -v -test.short

update:  ## Update all project dependencies
	go get -u ./...
	go mod tidy

update-releaser:  ## Update the goreleaser application
	brew update
	brew upgrade goreleaser

vet: ## Run the Go vet application
	go vet -v