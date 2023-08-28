GO ?= go

VERSION := $(shell git rev-parse --short HEAD)
BUILD_OUPUT = ./bin/
.PHONY: build
build:
	@echo "Running: go build version=$(VERSION)"
	@mkdir -p bin/ \
	&&  $(GO) build -ldflags="-X main.Version=$(VERSION)" ${BUILD_TAGS} -o $(BUILD_OUPUT) ./...