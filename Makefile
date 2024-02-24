build:
	go build -o ./.bin cmd/main/main.go

run: build
	./.bin

.PHONY: lint
lint:
	golangci-lint run -- config=.golangci.yaml

test:
	go test -race ./...