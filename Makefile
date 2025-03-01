SHELL=/bin/bash
BINARY_NAME=statd
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_ROOT := $(dir $(MAKEFILE_PATH))

build:
	cd $(PROJECT_ROOT) && mkdir -p $(PROJECT_ROOT)/bin && go build -o bin/$(BINARY_NAME) cmd/main.go

run:
	cd $(PROJECT_ROOT) && go run cmd/main.go

clean:
	cd $(PROJECT_ROOT) && go clean && rm -f bin/$(BINARY_NAME)

.PHONY: build run clean
