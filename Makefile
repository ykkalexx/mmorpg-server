build:
	@go build -o bin/go-mmorpg-server

run: build
	@./bin/go-mmorpg-server

test:
	@go test -v ./...