# Project and package definitions
PROJECT_NAME := "cyone"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

# Go environment variables
GOARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)

# Build metadata
BUILDDATE := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
CODEVERSION := "0.0.1"
CODEBUILDREVISION := $(shell git rev-parse HEAD)

# Phony targets to avoid conflicts with files of the same name
.PHONY: all dep build clean

# Default target to build the project
all: build

# Dependency management
dep: ## Get the dependencies
	@echo "  >  Getting dependencies..."
	@go mod download

# Build the project
build: dep
	@echo "  >  Building binary for $(GOOS)/$(GOARCH)..."
	GOARCH=$(GOARCH) GOOS=$(GOOS) BUILDDATE=$(BUILDDATE) CODEBUILDREVISION=$(CODEBUILDREVISION) go build -v -ldflags "\
		-X main.GOOS=$(GOOS) \
		-X main.GOARCH=$(GOARCH) \
		-X main.CODEVERSION=$(CODEVERSION) \
		-X main.CODEBUILDDATE=$(BUILDDATE) \
		-X main.CODEBUILDREVISION=$(CODEBUILDREVISION)" \
		-o $(PKG) cmd/cyone/main.go
	@mv $(PROJECT_NAME) "$(PROJECT_NAME)-$(GOOS)-$(GOARCH)"

# Clean up the previous build
clean: ## Remove previous build
	@echo "  >  Cleaning up previous build..."
	@-rm $(PROJECT_NAME) 2>/dev/null || true
