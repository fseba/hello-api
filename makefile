GO_VERSION :=1.18
TAG := $(shell git describe --abbrev=0 --tags --always)
HASH := $(shell git rev-parse HEAD)
DATE := $(shell date -u +"%Y-%m-%d.%H:%M:%SZ")
LDFLAGS := -w -X github.com/fseba/hello-api/handlers.hash=$(HASH) -X github.com/fseba/hello-api/handlers.tag=$(TAG) -X github.com/fseba/hello-api/handlers.date=$(DATE)

.PHONY: install-go init-go

setup: install-go init-go

build:
	go build -ldflags "$(LDFLAGS)" -o api cmd/main.go

test:
	go test -v ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | awk '{print (int($$3) <= 70) }'

report:
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $$(go fmt ./...)

#TODO: add MacOS support
install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linus-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linus-amd64.tar.gz
	rm go$(GO_VERSION).linus-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/urs/local/go/bin' >> $${HOME}/.zshrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.zshrc

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@2.6.0

static-check:
	golangci-lint run ./...
