BINARY_NAME=spwgen
MAIN_PATH=cmd/main.go
GOPATH=$(shell go env GOPATH)

.PHONY: all build clean install

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

clean:
	go clean
	rm -f $(BINARY_NAME)

install:
	go install $(MAIN_PATH)
	mv $(GOPATH)/bin/main $(GOPATH)/bin/$(BINARY_NAME)

test:
	go test ./... 