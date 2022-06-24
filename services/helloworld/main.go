package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/Xennis/go-hello-grpc/services/helloworld/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func main() {
	grpcPort, set := os.LookupEnv("HELLO_WORLD_GRPC_PORT")
	if !set {
		log.Fatalf("env %s not set", "HELLO_WORLD_GRPC_PORT")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (svc *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: "Echo: " + in.GetMessage(),
	}, nil
}
