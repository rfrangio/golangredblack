.PHONY: build test

GOCACHE ?= /tmp/gocache
BIN_DIR := bin
DEMO_BIN := $(BIN_DIR)/demo

build:
	mkdir -p $(BIN_DIR)
	GOCACHE=$(GOCACHE) go build -o $(DEMO_BIN) ./cmd/demo

test:
	GOCACHE=$(GOCACHE) go test ./...
