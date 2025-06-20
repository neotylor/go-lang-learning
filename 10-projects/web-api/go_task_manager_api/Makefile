# Makefile for Unix/macOS/Linux

# Makefile for Go Task Manager API

.PHONY: swag run build clean tidy

# Generate Swagger docs
swag:
	swag init -g main.go

# Run the project (with docs generation)
run: swag
	go run main.go

# Build the binary (with docs generation)
build: swag
	go build -o bin/task-api main.go

# Tidy dependencies
tidy:
	go mod tidy

# Clean compiled binaries and Swagger docs
clean:
	rm -rf bin/
	rm -rf docs/

# make run    # runs swag init and starts server
# make swag   # just regenerates swagger docs
# make build  # builds binary after updating docs
