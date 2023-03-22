# Set variables for the Makefile
GO_BINARY_NAME = "simple-server"
GO_BINARY_PATH = "./bin/$(GO_BINARY_NAME)"

LDFLAGS = "-s -w"

# Build the binary and run the tests
all: build test

# tidy the go.mod file
tidy:
	@echo "Tidying go.mod..."
	go mod tidy -v
	

# Build the binary
build: clean tidy
	@echo "Building $(GO_BINARY_NAME)..."
	go build -ldflags $(LDFLAGS) -o $(GO_BINARY_PATH) "cmd/$(GO_BINARY_NAME)/main.go"

# Run the binary
run: build
	@echo "Running $(GO_BINARY_NAME)..."
	./$(GO_BINARY_PATH)

# Run the tests
test: clean tidy
	@echo "Running tests..."
	go clean -testcache
	go test -v -cover ./...

# Clean the binary
clean:
	@echo "Cleaning up generated files..."
	rm -rfv bin