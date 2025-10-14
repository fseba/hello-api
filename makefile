GO_VERSION :=1.18

.PHONY: install-go init-go

setup: install-go init-go

build:
	go build -o api cmd/main.go

test:
	go test -v ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | awk '{print (int($$3) <= 70) }'

report:
	go tool cover -html=coverage.out -o cover.html

#TODO: add MacOS support
install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linus-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linus-amd64.tar.gz
	rm go$(GO_VERSION).linus-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/urs/local/go/bin' >> $${HOME}/.zshrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.zshrc

