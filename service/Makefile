
.PHONY: build

GOCMD:=go
GOVERSION:=go version
GOBUILD:=$(GOCMD) build
GOCLEAN:=$(GOCMD) clean
BINARY_NAME:=./build/
BINARY_UNIX:=$(BINARY_NAME)_unix

build:
	$(GOVERSION)
	which $(GOCMD)
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)
	rm -rf $(BINARY_UNIX)

all: clean build