.PHONY: default
default: generate build

.PHONY: generate
generate:
	go run . ./fixtures/index.json

.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o ./wasm/search.wasm ./wasm/main.go

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run
