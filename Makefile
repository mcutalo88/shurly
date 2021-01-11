SHELL:=/usr/bin/env bash

.PHONY: build
build:
	go build -o build/sherly

.PHONY: dev
dev:
	air

.PHONY: test
test:
	go test -race -covermode=atomic ./...

lint:
	golangci-lint run
