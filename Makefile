.PHONY: build test
build: 
	go build -v ./cmd/generator
test:
	go test -v -race -timeout 10s ./...

.DEFAULT_GOAL :=build 