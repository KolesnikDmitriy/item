.PHONY: version
version:
	go version
	protoc --version

.PHONY: generate
generate:
	protoc --go_out=pkg --go-grpc_out=pkg api/item.proto

.PHONY: run
run:
	go run ./cmd/item/main.go
