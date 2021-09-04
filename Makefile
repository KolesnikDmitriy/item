.PHONY: version
version:
	go version
	protoc --version

.PHONY: generate
generate:
	protoc --go_out=pkg --go_opt=paths=source_relative \
		--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
		api/item.proto

.PHONY: run
run:
	go run ./cmd/item/main.go

.PHONY: test
test:
	go test ./test/tests -count=1 -v
