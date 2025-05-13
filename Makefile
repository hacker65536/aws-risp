# AWS-RISP Makefile

# Variables
BINARY_NAME=aws-risp
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_VERSION=$(shell go version | awk '{print $$3}')
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${BUILD_DATE} -X github.com/hacker65536/aws-risp/cmd.version=${VERSION} -X github.com/hacker65536/aws-risp/cmd.commit=${COMMIT} -X github.com/hacker65536/aws-risp/cmd.date=${BUILD_DATE}"
PKG_LIST=$(shell go list ./... | grep -v /vendor/)

# Targets
.PHONY: all build clean test lint vet fmt help install uninstall release doc cover deps report

all: clean test build

# Build the application
build:
	@echo "Building ${BINARY_NAME} version ${VERSION} commit ${COMMIT}"
	go build ${LDFLAGS} -o ${BINARY_NAME} main.go

# Install the application
install:
	@echo "Installing ${BINARY_NAME}"
	go install ${LDFLAGS}

# Uninstall the application
uninstall:
	@echo "Uninstalling ${BINARY_NAME}"
	go clean -i

# Clean the build artifacts
clean:
	@echo "Cleaning build artifacts"
	go clean
	rm -f ${BINARY_NAME}
	rm -rf dist/

# Run tests
test:
	@echo "Running tests"
	go test -v ./...

# Run tests with coverage
cover:
	@echo "Running tests with coverage"
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

# Run benchmarks
bench:
	@echo "Running benchmarks"
	go test -bench=. ./...

# Run linting
lint:
	@echo "Running linter"
	golint ./...

# Run vet
vet:
	@echo "Running go vet"
	go vet ./...

# Format the code
fmt:
	@echo "Formatting code"
	go fmt ./...

# Create a release build for multiple platforms
release:
	@echo "Creating release build for version ${VERSION}"
	mkdir -p dist
	# Linux
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o dist/${BINARY_NAME}-linux-amd64 main.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o dist/${BINARY_NAME}-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o dist/${BINARY_NAME}-darwin-arm64 main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o dist/${BINARY_NAME}-windows-amd64.exe main.go
	@echo "Release builds created in ./dist/"

# Run the application with default arguments
run:
	@echo "Running ${BINARY_NAME}"
	go run ${LDFLAGS} main.go

# Show help
help:
	@echo "AWS-RISP Makefile Help"
	@echo ""
	@echo "Available targets:"
	@echo "  all        - Clean, test, and build"
	@echo "  build      - Build the application"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  cover      - Run tests with coverage report"
	@echo "  bench      - Run benchmarks"
	@echo "  lint       - Run linter"
	@echo "  vet        - Run go vet"
	@echo "  fmt        - Format code"
	@echo "  install    - Install the application"
	@echo "  uninstall  - Uninstall the application"
	@echo "  release    - Create release builds for multiple platforms"
	@echo "  run        - Run the application"
	@echo "  doc        - Generate documentation"
	@echo "  deps       - Install and update dependencies"
	@echo "  report     - Generate Go report card"
	@echo "  help       - Show this help"
	@echo ""
	@echo "Current variables:"
	@echo "  BINARY_NAME: ${BINARY_NAME}"
	@echo "  VERSION: ${VERSION}"
	@echo "  COMMIT: ${COMMIT}"
	@echo "  BUILD_DATE: ${BUILD_DATE}"
	@echo "  GO_VERSION: ${GO_VERSION}"

# Generate code documentation
doc:
	@echo "Generating documentation"
	godoc -http=:6060
	@echo "Documentation available at http://localhost:6060/pkg/github.com/hacker65536/aws-risp/"

# Install and update dependencies
deps:
	@echo "Installing dependencies"
	go mod tidy
	go mod download
	@echo "Dependencies updated"

# Generate a Go report card
report:
	@echo "Generating Go report card"
	@echo "Running go vet..."
	go vet ./...
	@echo "Running golint..."
	golint ./...
	@echo "Running gocyclo..."
	gocyclo -over 15 .
	@echo "Running ineffassign..."
	ineffassign ./...
