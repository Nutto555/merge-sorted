.DEFAULT_GOAL := help

help: ## list targets
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-8s %s\n", $$1, $$2}'

run: ## run the example
	go run ./cmd/merge

test: ## unit tests, verbose
	go test -race -v ./...

cover: ## coverage
	go test -cover ./...

lint: ## vet and formatting check
	go vet ./... && test -z "$$(gofmt -l .)"
