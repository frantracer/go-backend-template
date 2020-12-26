GOPATH_BIN=$(shell go env GOPATH)/bin

run:
	go run src/cmd/main.go

lint:
	$(GOPATH_BIN)/golangci-lint run

tools-dev: tool-linter

tool-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_BIN) v1.33.0
