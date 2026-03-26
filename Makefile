.PHONY: build test benchmark clean

GOCACHE ?= /tmp/gocache
BIN_DIR := bin
DEMO_BIN := $(BIN_DIR)/demo

build:
	mkdir -p $(BIN_DIR)
	GOCACHE=$(GOCACHE) go build -o $(DEMO_BIN) ./cmd/demo

test:
	GOCACHE=$(GOCACHE) go test ./...

benchmark: build
	/usr/bin/time ./$(DEMO_BIN)

clean:
	rm -rf $(BIN_DIR)
	find . -type f \( -name '*.o' -o -name '*.obj' \) -delete
