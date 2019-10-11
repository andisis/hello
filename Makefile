run: build
	@./bin/hello

test:
	@go fmt
	@go vet
	@go test

build:
	@go build -o bin/hello