SHELL := /bin/bash

.PHONY: all
all: deps test

.PHONY: generate-mocks
generate-mocks:
	@echo "===> Generating mocks"
	@mockery -dir=./types/slack -output=./types/slack/mocks -all

.PHONY: deps
deps:
	@echo "===> Fetching dependencies"
	@GO111MODULE=on go mod vendor

.PHONY: test
test:
	@echo "===> Running tests"
	@GO111MODULE=on go test -mod=vendor ./...
