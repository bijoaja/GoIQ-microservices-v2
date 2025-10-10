SERVICE_NAME=user-service
.PHONY: build run migrate

build:
	go build -o $(SERVICE_NAME) ./cmd/$(SERVICE_NAME)

run:
	go run ./cmd/$(SERVICE_NAME)

migrate:
	go run ./cmd/$(SERVICE_NAME) --migrate
