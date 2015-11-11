NAME = lsami
VERSION = $(shell cat VERSION)
REVISION = $(shell git rev-parse --short HEAD)

build:
	@go build -o build/$(NAME) -ldflags "-X main.Name=$(NAME) -X main.Version=$(VERSION) -X main.Revision=$(REVISION)"

clean:
	@go clean
	@rm -rf build

.PHONY: build clean
