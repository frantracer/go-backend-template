GOPATH_BIN=$(shell go env GOPATH)/bin

# Set GOROOT is only required for Github Actions image, it sets an old version of Go by default
export GOROOT=/usr/local/go

run:
	go run src/cmd/main.go

test:
	go test ./...

lint:
	$(GOPATH_BIN)/golangci-lint run

install-go:
	wget -c https://golang.org/dl/go1.15.6.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
	sudo ln -s /usr/local/go/bin/go /usr/bin/go

uninstall-go:
	sudo rm -rf /usr/local/go
	sudo rm -rf /usr/bin/go

setup-ci-env: setup-linter

setup-dev-env: setup-linter setup-act

setup-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.34.1/install.sh | sh -s -- -b $(GOPATH_BIN) v1.34.1

setup-act:
	curl https://raw.githubusercontent.com/nektos/act/v0.2.17/install.sh | sudo bash

act-push:
	act -P push