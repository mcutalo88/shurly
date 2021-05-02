SHELL:=/usr/bin/env bash

.PHONY: all
all: build lint test

.PHONY: build
build:
	go build -o build/sherly

.PHONY: dev
dev:
	air -c .air.toml

.PHONY: test
test:
	go test -race -covermode=atomic ./...

lint:
	golangci-lint run
