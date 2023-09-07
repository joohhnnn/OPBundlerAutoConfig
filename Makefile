# Variables
BINARY_NAME=OPBundlerAutoConfig
BUILD_DIR=build
SRC=./cmd/main.go

# Determine the operating system
OS := $(shell uname)

# Build the project
build:
	@echo "Building..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC)

# Install the binary into /usr/local/bin
install: build
	@echo "Installing..."
	sudo mv $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin
	chmod +x ./scripts/*.sh

# Clean up
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)


# Run all steps
all: build install

.PHONY: build install clean all
