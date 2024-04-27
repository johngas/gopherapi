build:
	@go build -o bin/gopherapi

run: build
	@./bin/gopherapi

test:
	@go test -v ./...