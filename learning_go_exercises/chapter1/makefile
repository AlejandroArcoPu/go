.DEFAULT_GOAL := build

.PHONY: fmt build

fmt:
	go fmt

build: fmt 
	go build .

clean: build
	rm -rf ./hello_world
