# Makefile
# Default target
.PHONY: all
all: build

# Build the project (optional)
.PHONY: build
build:
	go build ./...

# Run all tests in the tests/ folder
.PHONY: test
test:
	go test -v ./tests

# Run tests with race detection
.PHONY: test-race
test-race:
	go test -race ./...

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint (if using golangci-lint)
.PHONY: lint
lint:
	golangci-lint run ./...

# Clean up
.PHONY: clean
clean:
	go clean
