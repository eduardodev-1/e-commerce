# Makefile for Go project

# Variables
APP_NAME := map5
SOURCE_DIR := ./cmd/api
BUILD_DIR := ./build

# Determine the operating system
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Windows_NT)
	# Windows
	EXECUTABLE := $(APP_NAME).exe
	CLEAN_COMMAND := del /Q $(BUILD_DIR)\*
else
	# Linux or macOS
	EXECUTABLE := $(APP_NAME)
	CLEAN_COMMAND := rm -rf $(BUILD_DIR)/*
endif

# Default target
.PHONY: all
all:
	@echo "operating system: $(UNAME_S)"


# Run the application
.PHONY: run
run:
	go run $(SOURCE_DIR)/*.go

# Build the application
.PHONY: build
build: clean
	@mkdir -p $(BUILD_DIR)
	@echo "Building Linux binary"
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SOURCE_DIR)/*.go
	@echo "Build complete."

	@echo "Building Windows executable"
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME).exe $(SOURCE_DIR)/*.go
	@echo "Build complete."

# Build and run the application
.PHONY: buildandrun
buildandrun: build
	@echo "OS: $(UNAME_S)"
	@echo "Running $(EXECUTABLE)..."
	$(BUILD_DIR)/$(EXECUTABLE)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(CLEAN_COMMAND)
	@echo "Clean complete."
