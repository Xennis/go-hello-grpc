
setup:
	@echo "Install tools"
	go install github.com/bufbuild/buf/cmd/buf
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	@echo "Generate rpc"
	buf generate

run:
	@echo "Run helloworld service"
	HELLO_WORLD_GRPC_PORT=5001 go run services/helloworld/main.go
