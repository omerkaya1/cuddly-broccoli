PROJECT_NAME=cuddly-broccoli

# Build targets
.PHONY: build test
# Code formatting, linting and UTs
.PHONY: fmt vet lint
# Supplementary targets
.PHONY: help clean

vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v sandbox)

fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v ^vendor/ | tee /dev/stderr)" || \
		(echo >&2 "+ please format Go code with 'gofmt -s'" && false)

check-lint: ## Run a check to verify whether golangci-lint is installed locally
	@which golangci-lint || (GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint)

lint: check-lint ## Run selected linters
	@echo "+ $@"
	@golangci-lint run -v --timeout 3m \
		--issues-exit-code=0 \
		--print-issued-lines=false \
		--enable=gocognit \
		--enable=gocritic \
		--enable=prealloc \
		--enable=unparam \
		--enable=nakedret \
		--enable=exportloopref \
		--enable=deadcode  \
		--enable=unused  \
		--enable=gocyclo \
		--enable=revive \
		--enable=varcheck \
        --enable=structcheck \
        --enable=govet \
        --enable=errcheck \
        --enable=dupl \
        --enable=lll \
        --enable=ineffassign \
        --enable=unconvert \
        --enable=goconst \
        --enable=gosec \
        --enable=megacheck \
		./...

build: lint ## Builds the project's binary
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ./$(PROJECT_NAME) ./cmd/cuddly-broccoli/.

test: lint ## Runs UTs on the project files
	@go test ./... -count=1 -race

coverage: ## Runs UTs and reports coverage
	@go test ./... -count=1 -race -coverprofile=coverage.txt && go tool cover -html=coverage.txt

clean: ## Remove all temporary files
	rm -f $(PROJECT_NAME) *.txt *.exe

help: ## Print this help and exit
	@grep -E '^[a-zA-Z_\-\/]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
