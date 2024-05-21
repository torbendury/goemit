.PHONY: all test example

all: test

test:
	go test -v ./... -race

example:
	go run example/main.go
