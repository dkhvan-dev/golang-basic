.PHONY: build
build:
	go run -v ./cmd

.DEFAULT_GOAL := build