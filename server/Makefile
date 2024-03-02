# Simple Makefile for a Go project

# Build the application
all: build

build:
	@tailwindcss -i cmd/web/css/_app.css -o cmd/web/css/styles.css
	@echo "Building..."
	@templ generate
	@go build -o main cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

tailwind:
	@tailwindcss -i cmd/web/css/_app.css -o cmd/web/css/styles.css

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@templ generate
	@tailwindcss -i cmd/web/css/_app.css -o cmd/web/css/styles.css
	@air

.PHONY: all build run test clean
