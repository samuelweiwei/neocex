APP_NAME=neocex
BIN_DIR=bin
ENTRY=serve.go

.PHONY: help setup build run clean fmt vet test tidy

help:
	@echo "Targets:"
	@echo "  make setup  - prepare local folders"
	@echo "  make build  - build binary from $(ENTRY)"
	@echo "  make run    - run service locally"
	@echo "  make fmt    - format all go files"
	@echo "  make vet    - run go vet"
	@echo "  make test   - run go tests"
	@echo "  make tidy   - tidy go modules"
	@echo "  make clean  - remove build artifacts"

setup:
	@mkdir -p $(BIN_DIR)

build: setup
	go build -o $(BIN_DIR)/$(APP_NAME) ./$(ENTRY)

run:
	go run ./$(ENTRY)

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

vet:
	go vet ./...

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)
