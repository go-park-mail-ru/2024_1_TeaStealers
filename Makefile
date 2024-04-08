OS := $(shell uname -s)

ifneq ("$(wildcard .env)","")
include .env
endif

ifeq ($(OS), Linux)
	DOCKER_COMPOSE := docker compose
endif
ifeq ($(OS), Darwin)
	DOCKER_COMPOSE := docker-compose
endif

build_:
	go build -o ./.bin cmd/main/main.go

run: build_
	./.bin

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yaml

test:
	go test -race ./...

migrate-lib:
	go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/

create-migration:
	migrate create -dir migrations -ext sql -seq $(TABLE_NAME)

migrate-up:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

dev-compose-up:
	$(DOCKER_COMPOSE) -f "dev-docker-compose.yaml" up -d

dev-compose-down:
	$(DOCKER_COMPOSE) -f "dev-docker-compose.yaml" down

coverage:
	go test -json ./... -coverprofile coverprofile_.tmp -coverpkg=./... ; \
	cat coverprofile_.tmp |grep -v auth.go| grep -v interfaces.go | grep -v docs.go| grep -v cors.go| grep -v transaction.go| grep -v main.go > coverprofile.tmp ; \
	rm coverprofile_.tmp ; \
	go tool cover -html coverprofile.tmp ; \
	go tool cover -func coverprofile.tmp

swagger:
	swag init -g cmd/main/main.go

