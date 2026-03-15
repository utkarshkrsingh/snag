SOURCE_DIR := ./cmd/snag/
BINARY_NAME := snag
TARGET_DIR := ./bin
TEST_COVERAGE_DIR := ./coverage
PKG_DEPENDENCIES := github.com/fatih/color github.com/fsnotify/fsnotify github.com/spf13/cobra github.com/spf13/viper

all: deps build run

deps:
	## Downloading all package dependencies
	@go mod tidy
	@go get $(PKG_DEPENDENCIES)

build:
	## Building the binary
	@go build -o $(TARGET_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

run:
	## Running the binary
	./bin/snag

test:
	@go test -v ./cmd/snag

test-coverage:
	mkdir -p $(TEST_COVERAGE_DIR)
	@go test ./... -coverprofile=$(TEST_COVERAGE_DIR)/coverage.out
	@go tool cover -html=coverage/coverage.out -o coverage/index.html


clean:
	## Clearing binary
	@$(RM) -f $(TARGET_DIR)/$(BINARY_NAME)

.PHONY: all deps build run test test-coverage clean
