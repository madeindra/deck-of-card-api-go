.PHONY: install run build

install:
	go mod download

run:
	go run ./main.go

build:
	go build -o main

test:
	mkdir -p ./coverage && \
		go test -v -coverprofile=./coverage/coverage.out -covermode=atomic ./...

cover: test
	go tool cover -func=./coverage/coverage.out &&\
		go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html	