.PHONY: all format test bench
all: format test

format:
	go fmt ./...

test:
	go test -v ./...

bench:
	go test -v -bench=. ./...
