SOURCE_DIR := ./cmd/snag/
BINARY_NAME := snag
TARGET_DIR := ./bin
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

clean:
	## Clearing binary
	@$(RM) -f $(TARGET_DIR)/$(BINARY_NAME)

.PHONY: all deps build run clean
