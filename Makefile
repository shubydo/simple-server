# Set variables for the Makefile
GO_BINARY_NAME = "simple-server"
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
	go build -ldflags $(LDFLAGS) -o $(GO_BINARY_NAME) main.go

# Run the binary
run: build
	@echo "Running $(GO_BINARY_NAME)..."
	./$(GO_BINARY_NAME)

# Run the tests
test: clean tidy
	@echo "Running tests..."
	go clean -testcache
	go test -v -cover ./...

# Clean the binary
clean:
	@echo "Cleaning $(GO_BINARY_NAME)..."
	rm -fv $(GO_BINARY_NAME)