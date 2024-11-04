SOURCE_DIR := .
BINARY_NAME := runsync
TARGET_DIR := /usr/bin
PKG_DEPENDENCIES := github.com/fatih/color github.com/fsnotify/fsnotify

all: build

build: deps
	@go build -o $(SOURCE_DIR)/$(BINARY_NAME) .

deps:
	@go mod tidy
	@go get $(PKG_DEPENDENCIES)

install:
	@mkdir -p $(TARGET_DIR)
	@cp $(SOURCE_DIR)/$(BINARY_NAME) $(TARGET_DIR)
	@chmod +x $(TARGET_DIR)/$(BINARY_NAME)

clean:
	@$(RM) -f $(BINARY_NAME)

uninstall:
	@$(RM) -f $(TARGET_DIR)/$(BINARY_NAME)

.PHONY: all build deps install clean uninstall
