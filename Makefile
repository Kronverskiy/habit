#!/usr/bin/make -f

BIN_DIR = ./bin
BIN_NAME = habit
SRC_DIR = ./cmd/habit

.PHONY: all build oapi-gen-install generate-api

all: build

# Build habit service
build:
	@echo "⇒ Building $(BIN_NAME)..."
	@mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 go build -v -o $(BIN_DIR)/$(BIN_NAME) $(SRC_DIR)
	@echo "⇒ Building $(BIN_NAME) complete!"

# Install binary openapi generator
oapi-gen-install:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# Generate api from openapi spec
generate-api:
	go generate ./api/...
