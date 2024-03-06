OS := $(shell uname -s)

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
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

dev-compose-up:
	$(DOCKER_COMPOSE) -f "dev-docker-compose.yaml" up -d

dev-compose-down:
	$(DOCKER_COMPOSE) -f "dev-docker-compose.yaml" down

swagger:
	swag init -g cmd/main/main.go
