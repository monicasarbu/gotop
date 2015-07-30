BIN_PATH?=/usr/bin

GOFILES = $(shell find . -type f -name '*.go')
gotop: $(GOFILES)
	go build

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm gotop || true
