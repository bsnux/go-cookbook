GONAME=$(shell basename "$(PWD)")
GOARCH=amd64

LINUXBIN=/usr/local/bin
DARWINBIN=/usr/local/bin

GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

test:
	go test -v

linux:
	@echo "Building binary for linux..."
	GOOS=linux GOARCH=amd64 go build -o bin/$(GONAME)-linux-amd64

darwin:
	@echo "Building binary for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o bin/$(GONAME)-darwin-amd64

windows:
	@echo "Building binary for Windows..."
	GOOS=windows GOARCH=amd64 go build -o bin/$(GONAME)-windows-amd64

clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

install-linux:
	@echo "Copying binary to $(LINUXBIN) ..."
	cp bin/$(GONAME)-linux-amd64 $(LINUXBIN)/$(GONAME)

install-darwin:
	@echo "Copying binary to $(DARWINBIN) ..."
	cp bin/$(GONAME)-darwin-amd64 $(DARWINBIN)/$(GONAME)

all: linux darwin windows

.PHONY: test linux darwin windows clean all
