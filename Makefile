# Binary name
BINARY_NAME=rewardmanager

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

# Main package path
MAIN_PATH=./cmd/rewardmanager

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

# Run the application
run:
	$(GORUN) $(MAIN_PATH)

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Build and run
all: build run

# Help command
help:
	@echo "Available commands:"
	@echo "  make build  - Build the application"
	@echo "  make run    - Run the application"
	@echo "  make clean  - Clean build files"
	@echo "  make test   - Run tests"
	@echo "  make all    - Build and run the application"

.PHONY: build run clean test all help 