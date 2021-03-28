GOPATH_BIN=$(shell go env GOPATH)/bin

# Set GOROOT is only required for Github Actions image, it sets an old version of Go by default
export GOROOT=/usr/local/go

run:
	go run src/cmd/main.go

test:
	go test ./...

lint:
	$(GOPATH_BIN)/golangci-lint run

mock:
	$(GOPATH_BIN)/moq -out src/infrastructure/http/zmock_app_test.go -pkg http_test src/application/handlers \
InsertTaskHandler FindTasksHandler

install-go:
	wget -c https://golang.org/dl/go1.16.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
	sudo ln -s /usr/local/go/bin/go /usr/bin/go

uninstall-go:
	sudo rm -rf /usr/local/go
	sudo rm -rf /usr/bin/go

setup-ci-env: setup-linter

setup-dev-env: setup-linter setup-act setup-mocker

setup-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.34.1/install.sh | sh -s -- -b $(GOPATH_BIN) v1.34.1

setup-act:
	curl https://raw.githubusercontent.com/nektos/act/v0.2.17/install.sh | sudo bash

setup-mocker:
	go install github.com/golang/mock/mockgen@v1.5.0

act-push:
	act -P push