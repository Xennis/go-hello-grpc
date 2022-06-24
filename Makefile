
setup:
	@echo "Install tools"
	go install github.com/bufbuild/buf/cmd/buf \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	@echo "Generate rpc"
	buf generate

lint:
	@echo "Lint proto"
	buf lint

run-helloworld:
	@echo "Run helloworld service"
	HELLO_WORLD_GRPC_PORT=5001 go run services/helloworld/main.go

run-restgateway:
	@echo "Run restgateway service"
	REST_GATEWAY_HTTP_PORT=6001 REST_GATEWAY_GRPC_ADDR=:5001 go run services/restgateway/main.go
