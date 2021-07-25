Go := go
Golangci_Lint := golangci-lint

.PHONY: build
build:
	$(Go) build ./cmd/...

.PHONY: install
install:
	$(Go) install ./cmd/...

.PHONY: test
test: check

.PHONY: check
check:
	$(Go) test ./...

.PHONY: lint
lint:
	$(Go) mod tidy
	$(Golangci_Lint) run ./...
