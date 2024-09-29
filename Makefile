# Variables
GO=go
EXECUTABLE_NAME=mcjsonwriter

# Detect OS and set executable extension
ifeq ($(OS),Windows_NT)
	EXECUTABLE_NAME := $(EXECUTABLE_NAME).exe
endif

# Default target
all: build

# Build the executable
build:
	$(GO) build -o $(EXECUTABLE_NAME) main.go item_models.go json_writer.go

# Clean up
clean:
	$(GO) clean
	rm -f $(EXECUTABLE_NAME)

.PHONY: all build clean