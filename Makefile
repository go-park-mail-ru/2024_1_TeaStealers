build:
	go build -o ./.bin cmd/main/main.go

run: build
	./.bin

lint:
	golangci-lint run

test:
	go test ./...