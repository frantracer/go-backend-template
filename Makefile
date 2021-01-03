GOPATH_BIN=$(shell go env GOPATH)/bin

run:
	go run src/cmd/main.go

test:
	go test ./...

lint:
	$(GOPATH_BIN)/golangci-lint run

install-go:
	wget -c https://golang.org/dl/go1.14.13.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
	ln -s /usr/local/go/bin/go /usr/bin/go

uninstall-go:
	rm -rf /usr/local/go
	rm -rf /usr/bin/go

setup-ci-env: setup-linter

setup-dev-env: setup-linter

setup-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_BIN) v1.34.1

act-push:
	act -P push