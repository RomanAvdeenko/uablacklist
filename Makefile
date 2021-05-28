.PHONY: build test
build: 
	go build -v ./cmd/uablacklist
test:
	go test -v -race -timeout 10s ./...

.DEFAULT_GOAL :=build 