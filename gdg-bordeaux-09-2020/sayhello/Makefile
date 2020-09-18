#!/usr/bin/make
.PHONY: test

default: build

go-generate:
	go generate ./...

test: go-generate
	go test -v -test.count 1 ./...

build: go-generate
	go install ./...