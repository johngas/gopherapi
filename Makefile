build:
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

run: build
	@./bin/$(APP_NAME)

test:
	@go test -v ./...