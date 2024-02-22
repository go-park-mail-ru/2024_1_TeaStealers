build:
	go build -o ./.bin cmd/main/main.go

run: build
	./.bin