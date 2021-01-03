GOPATH_BIN=$(shell go env GOPATH)/bin

run:
	go run src/cmd/main.go

test:
	go test ./...

lint:
	$(GOPATH_BIN)/golangci-lint run

install-go:
	wget -c https://golang.org/dl/go1.15.6.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
	ln -s /usr/local/go/bin/go /usr/bin/go

uninstall-go:
	rm -rf /usr/local/go
	rm -rf /usr/bin/go

setup-ci-env: setup-linter

setup-dev-env: setup-linter

setup-linter:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.34.1

act-push:
	act -P ubuntu-latest=nektos/act-environments-ubuntu:18.04 push