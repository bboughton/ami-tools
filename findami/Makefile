NAME = findami
VERSION = $(shell cat VERSION)
REVISION = $(shell git rev-parse --short HEAD)

GOCMD = go

build:
	@$(GOCMD) build -ldflags "-X main.Name=$(NAME) -X main.Version=$(VERSION) -X main.Revision=$(REVISION)"

clean:
	@$(GOCMD) clean
	@rm -rf build

install:
	@$(GOCMD) install
	@cp share/man/man1/$(NAME).1 $(GOPATH)/share/man/man1/$(NAME).1

uninstall:
	@$(RM) $(GOPATH)/bin/$(NAME)
	@$(RM) $(GOPATH)/share/man/man1/$(NAME).1

.PHONY: build clean install
