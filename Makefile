#!/usr/bin/make -f

BIN_DIR = ./bin
BIN_NAME = habit
SRC_DIR = ./cmd/habit

.PHONY: all build

all: build

build:
	@echo "⇒ Building $(BIN_NAME)..."
	@mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 go build -v -o $(BIN_DIR)/$(BIN_NAME) $(SRC_DIR)
	@echo "⇒ Building $(BIN_NAME) complete!"
