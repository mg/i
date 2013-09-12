all: install

build:
	go build ./...

install:
	go install ./...

test: 
	go test ./...
