SHELL := /bin/bash

.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## go run
	go run .

.PHONY: generate
generate: ## generate theme files
	go run . generate

.PHONY: fmt-check
fmt-check: ## ensure gofmt has been run
	@files=$$(gofmt -l .); \
	if [ -n "$$files" ]; then \
		echo "gofmt required for:"; \
		echo "$$files"; \
		exit 1; \
	fi

.PHONY: vet
vet: ## go vet static checks
	go vet ./...

.PHONY: test
test: ## go test
	go test ./...

.PHONY: ci
ci: ## run checks used by CI
ci: fmt-check vet test

.PHONY: copygen
copygen: ## copies generated themes to ~/.config/zed/themes
	go run . generate
	mkdir -p ~/.config/zed/themes
	cp themes/*.json ~/.config/zed/themes/
