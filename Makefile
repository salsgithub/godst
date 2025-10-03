define print-target
	@echo
	@printf "\033[31m*\033[0m Executing target: \033[31m$@\033[0m\n"
endef

.DEFAULT_GOAL := all

.PHONY: all
all: audit tidy test

.PHONY: help
help: ## List targets
	@printf "\033[31m[ Targets ]\033[0m\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[31m%-15s\033[0m %s\n", $$1, $$2}'


.PHONY: audit
audit: ## Audit for vulnerabilities
	$(call print-target)
	go mod tidy -diff
	go mod verify
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: tidy
tidy: ## Tidy dependencies and format
	$(call print-target)
	go mod tidy -v
	go fmt ./...

.PHONY: test
test: ## Run tests
	mkdir -p testresults
	go tool gotestsum --junitfile testresults/unit-tests.xml -- -race -covermode=atomic -coverprofile=testresults/cover.out -v ./...
	go tool cover -html=testresults/cover.out -o testresults/coverage.html

.PHONY: bench
bench: ## Run benchmark tests
	go test -run=NO_TEST -bench . -benchmem -benchtime 1s ./...