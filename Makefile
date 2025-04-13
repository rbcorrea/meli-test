.PHONY: build run test lint up down

build:
	go build -o bin/app ./cmd/api/main.go

run:
	go run ./cmd/api/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

up:
	docker-compose up --build -d

down:
	docker-compose down
