.PHONY: all test

all: test

test:
	go test -v ./... -race