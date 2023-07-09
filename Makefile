.PHONY: build clean

GIT_ROOT := $(shell git rev-parse --show-toplevel)

test:
	go test ./internal/app

build:
	go build -ldflags="-s -w" -o $(GIT_ROOT)/bin/tfg $(GIT_ROOT)/main.go

clean:
	rm -rf ./bin