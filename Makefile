BINARY_NAME=runsync
SOURCE_FILES=main.go
PKG_DEPENDENCIES=github.com/fatih/color v1.18.0 github.com/fsnotify/fsnotify v1.8.0

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) $(SOURCE_FILES)

# Run the application
run: build
	./$(BINARY_NAME)

# Test the application
test:
	go test ./...

# Install dependencies
deps:
	go mod tidy
	go get $(PKG_DEPENDENCIES)

# Clean the build
clean:
	rm -f $(BINARY_NAME)

.PHONY: all build run test deps clean
