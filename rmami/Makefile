NAME = rmami
VERSION = $(shell cat VERSION)
REVISION = $(shell git rev-parse --short HEAD)

GOCMD=go
LDFLAGS=-ldflags "-X main.Name=$(NAME) -X main.Version=$(VERSION) -X main.Revision=$(REVISION)"

build:
	@$(GOCMD) build $(LDFLAGS)

clean:
	@$(GOCMD) clean

install:
	@$(GOCMD) install $(LDFLAGS)

uninstall:
	@$(RM) $(GOPATH)/bin/$(NAME)

.PHONY: build clean install uninstall
