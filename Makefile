.PHONY: install run build

install:
	go mod download

run:
	go run ./main.go

build:
	go build -o main