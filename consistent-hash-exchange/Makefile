SHELL := /bin/sh

.PHONY: dep
dep:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	@rm -rf bin/
	@mkdir bin/
	@cd bin && go build -o . ../cmd/setup-queue
	@cd bin && go build -o . ../cmd/setup-consistent-hash-exchange
	@cd bin && go build -o . ../cmd/producer
	@cd bin && go build -o . ../cmd/consumer
