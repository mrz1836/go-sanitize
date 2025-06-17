# Default repository domain name
ifndef GIT_DOMAIN
       override GIT_DOMAIN := github.com
endif

# Set if defined (alias variable for ease of use)
ifdef branch
       override REPO_BRANCH := $(branch)
       export REPO_BRANCH
endif

# Do we have git available?
HAS_GIT := $(shell command -v git 2> /dev/null)

ifdef HAS_GIT
# Do we have a repo?
HAS_REPO := $(shell git rev-parse --is-inside-work-tree 2> /dev/null)
ifdef HAS_REPO
# Automatically detect the repo owner and repo name (for local use with Git)
               REPO_NAME := $(shell basename "$(shell git rev-parse --show-toplevel 2> /dev/null)")
               OWNER := $(shell git config --get remote.origin.url | sed 's/git@$(GIT_DOMAIN)://g' | sed 's/\/$$(REPO_NAME).git//g')
               REPO_OWNER := $(shell echo $(OWNER) | tr A-Z a-z)
               VERSION_SHORT := $(shell git describe --tags --always --abbrev=0)
               export REPO_NAME
               export REPO_OWNER
               export VERSION_SHORT
       endif
endif

# Set the distribution folder
ifndef DISTRIBUTIONS_DIR
       override DISTRIBUTIONS_DIR := ./dist
endif
export DISTRIBUTIONS_DIR

.PHONY: citation
citation: ## Update version in CITATION.cff (citation version=X.Y.Z)
	@echo "updating CITATION.cff version..."
	@if [ -z "$(version)" ]; then \
		echo "Error: 'version' variable is not set. Please set the 'version' variable before running this target."; \
		exit 1; \
	fi
	@if [ "$(shell uname)" = "Darwin" ]; then \
		sed -i '' -e 's/^version: ".*"/version: "$(version)"/' CITATION.cff; \
	else \
		sed -i -e 's/^version: ".*"/version: "$(version)"/' CITATION.cff; \
	fi

.PHONY: diff
diff: ## Show the git diff
	$(call print-target)
	git diff --exit-code
	RES=$$(git status --porcelain) ; if [ -n "$$RES" ]; then echo $$RES && exit 1 ; fi

.PHONY: help
help: ## Show this help message
	@grep -Eh '^(.+):\ ##\ (.+)' ${MAKEFILE_LIST} | sort | column -t -c 2 -s ':#'

.PHONY: update-readme
update-readme: ## Update the README.md with the make commands
	@echo "updating makefile commands in the README.md..."
	@TMP=$$(mktemp) && \
	make --no-print-directory -s help | grep -v '^$$' > $$TMP && \
	awk -v hf=$$TMP 'BEGIN{skip=0} /<!-- make-help-start -->/{print; print "```text"; while((getline line < hf)>0) print line; print "```"; skip=1; next} /<!-- make-help-end -->/{print; skip=0; next} !skip{print}' README.md > README.md.tmp && \
	mv README.md.tmp README.md && \
	rm $$TMP

.PHONY: install-releaser
install-releaser: ## Install the GoReleaser application
	@echo "installing GoReleaser..."
	@curl -sSfL https://install.goreleaser.com/github.com/goreleaser/goreleaser@latest | sh

.PHONY: release
release:: ## Full production release (creates release in GitHub)
	@echo "releasing..."
	@test -n "$(github_token)"
	@export GITHUB_TOKEN=$(github_token) && goreleaser --rm-dist

.PHONY: release-test
release-test: ## Full production test release (everything except deploy)
	@echo "creating a release test..."
	@goreleaser --skip-publish --rm-dist

.PHONY: release-snap
release-snap: ## Test the full release (build binaries)
	@echo "creating a release snapshot..."
	@goreleaser --snapshot --skip-publish --rm-dist

.PHONY: tag
tag: ## Generate a new tag and push (tag version=0.0.0)
	@echo "creating new tag..."
	@test $(version)
	@git tag -a v$(version) -m "Pending full release..."
	@git push origin v$(version)
	@git fetch --tags -f

.PHONY: tag-remove
tag-remove: ## Remove a tag if found (tag-remove version=0.0.0)
	@echo "removing tag..."
	@test $(version)
	@git tag -d v$(version)
	@git push --delete origin v$(version)
	@git fetch --tags

.PHONY: tag-update
tag-update: ## Update an existing tag to current commit (tag-update version=0.0.0)
	@echo "updating tag to new commit..."
	@test $(version)
	@git push --force origin HEAD:refs/tags/v$(version)
	@git fetch --tags -f

.PHONY: update-releaser
update-releaser:  ## Update the goreleaser application
	@echo "updating GoReleaser application..."
	@$(MAKE) install-releaser
