SHELL := /bin/sh
MIGRATION_PATH := ./migrations/

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

.PHONY: migrate
migrate:
	@migrate -path $(MIGRATION_PATH) \
		-database "mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@$(MYSQL_PROTOCOL)($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)" up

.PHONY: migrate-rollback
migrate-rollback:
	@migrate -path $(MIGRATION_PATH) \
		-database "mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@$(MYSQL_PROTOCOL)($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)" down 1

.PHONY: mocks
mocks:
	mockery --name=JobService --with-expecter

.PHONY: build
build:
	@rm -rf bin/
	@mkdir bin/
	@cd bin && go build -o . ../cmd/workerd

.PHONY: run
run:
	./bin/workerd

