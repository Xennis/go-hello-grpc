package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pbHelloworld "github.com/Xennis/go-hello-grpc/services/helloworld/proto"
	pb "github.com/Xennis/go-hello-grpc/services/restgateway/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	hwc pbHelloworld.HelloWorldServiceClient
}

func (svc *server) Echo(ctx context.Context, req *pbHelloworld.EchoRequest) (*pbHelloworld.EchoResponse, error) {
	return svc.hwc.Echo(ctx, req)
}

func main() {
	ctx := context.Background()
	httpPort, set := os.LookupEnv("REST_GATEWAY_HTTP_PORT")
	if !set {
		log.Fatalf("env %s not set", "REST_GATEWAY_HTTP_PORT")
	}
	grpcAddr, set := os.LookupEnv("REST_GATEWAY_GRPC_ADDR")
	if !set {
		log.Fatalf("env %s not set", "REST_GATEWAY_GRPC_ADDR")
	}

	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial grpc: %v", err)
	}

	mux := runtime.NewServeMux()
	svc := &server{
		hwc: pbHelloworld.NewHelloWorldServiceClient(conn),
	}
	if err := pb.RegisterApiGatewayServiceHandlerServer(ctx, mux, svc); err != nil {
		log.Fatalf("failed to register handler: %v", err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), mux); err != nil {
		log.Fatalf("failed to serve http: %v", err)
	}
}
