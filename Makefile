.PHONY: version
version:
	go version
	protoc --version

.PHONY: generate
generate:
	protoc -I ./ \
		--go_out ./pkg --go_opt paths=source_relative \
    	--go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./pkg \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./pkg \
		--openapiv2_opt logtostderr=true \
		api/item.proto

.PHONY: run
run:
	go run ./cmd/item/main.go

.PHONY: test
test:
	go test ./test/tests -count=1

.PHONY: test-ci
test-ci:
	go test ./test/tests -count=1 -v -json
