NAME = findami
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
	@cp share/man/man1/$(NAME).1 $(GOPATH)/share/man/man1/$(NAME).1

uninstall:
	@$(RM) $(GOPATH)/bin/$(NAME)
	@$(RM) $(GOPATH)/share/man/man1/$(NAME).1

.PHONY: build clean install uninstall
