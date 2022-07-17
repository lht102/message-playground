SHELL := /bin/sh

-include .env.$(ENV)
export

.PHONY: all
all: lint test build run

.PHONY: dep
dep:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./... -v -race -cover -coverprofile=c.out -covermode=atomic

.PHONY: migrate-create
migrate-create:
	go run -mod=mod cmd/migrate/main.go $(name)

.PHONY: mocks
mocks:
	mockery --name=Service --with-expecter

.PHONY: build
build:
	@rm -rf bin/
	@mkdir bin/
	@cd bin && go build -o . ../cmd/workerd

.PHONY: run
run:
	./bin/workerd

