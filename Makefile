# Set variables for the Makefile
GO_BINARY_NAME = simple-server
GO_BINARY_PATH = ./bin/$(GO_BINARY_NAME)
GOLANGCI_LINT_VERSION = v1.58.1
LDFLAGS = "-s -w"

# Use arguments passed to the make command
ARGS = $(filter-out $@,$(MAKECMDGOALS))

# Build the binary and run the tests
all: build test

.PHONY: all build run test test-coverage clean tidy lint super-lint golangci-lint
super-lint:
	@echo "Linting all code EXCEPT GO with super-lint..."
	docker run --rm --name super-lint \
	    --env-file ".github/super-linter.env" \
		-v $(shell pwd):/tmp/lint \
		-e RUN_LOCAL=true \
		github/super-linter:slim-v5

golangci-lint:
	@echo "Linting all Go code with golangci-lint"
	docker run --rm --name golangci-lint  \
		-v $(shell pwd):/app \
		-v ~/.cache/golangci-lint/$(GOLANGCI_LINT_VERSION)/root/.cache  \
		-w /app \
		golangci/golangci-lint:$(GOLANGCI_LINT_VERSION) \
		golangci-lint run -v

lint: super-lint golangci-lint
	@echo "Linting..."

# tidy the go.mod file
tidy:
	@echo "Tidying go.mod..."
	go mod tidy -v
	

# Build the binary
build: clean tidy
	@echo "Building $(GO_BINARY_NAME)..."
	go build -ldflags=$(LDFLAGS) -o $(GO_BINARY_PATH) 

# Run the binary
run: build
	@echo "Running $(GO_BINARY_NAME)..."
	./$(GO_BINARY_PATH)

# Run the tests
test: clean tidy
	@echo "Running tests..."
	go clean -testcache
	go test -v -cover ./...


# Run the tests with coverage
test-coverage: clean tidy
	@echo "Running tests with coverage..."
	go clean -testcache
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean the binary
clean:
	@echo "Cleaning up generated files..."
	rm -rfv coverage.* $(GO_BINARY_PATH) tmp

# Install CI tools
# TODO: move to script or a dedicated `tools` package
.PHONY: tools
tools:
	@echo "Installing CI tools..."
	@echo "Installing cosmtrek/air..."
	go install github.com/cosmtrek/air@latest

	@echo "Installing golangci/golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
