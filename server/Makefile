# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@tailwindcss -i public/_app.css -o public/styles.css
	@templ generate
	@go build -o main .

# Run the application
run:
	go run .

tailwind:
	@tailwindcss -i public/_app.css -o public/styles.css

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@templ generate --watch & \
	@tailwindcss -i public/_app.css -o public/styles.css --watch & \
	air
 
.PHONY: all build run test clean
