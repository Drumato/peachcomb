all: format test

format:
	go fmt ./...

test:
	go test ./...

