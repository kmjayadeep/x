BINARY := x
PACKAGE := ./cmd/x
GO ?= go

.DEFAULT_GOAL := build

.PHONY: build check clean fmt fmt-check install nix-build test tidy vendor vet

build:
	$(GO) build -trimpath -o $(BINARY) $(PACKAGE)

test:
	$(GO) test -race ./...

vet:
	$(GO) vet ./...

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

fmt-check:
	@test -z "$$(gofmt -l $$(find . -name '*.go' -not -path './vendor/*'))" || \
		(echo "Go files need formatting; run 'make fmt'" && exit 1)

check: fmt-check vet test

tidy:
	$(GO) mod tidy

vendor: tidy
	$(GO) mod vendor

install:
	$(GO) install $(PACKAGE)

nix-build:
	nix build 'path:.#' --no-link

clean:
	$(GO) clean
	rm -f $(BINARY) coverage.out
