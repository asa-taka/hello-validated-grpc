package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/asa-taka/hello-validated-grpc/api"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type greetingServer struct {
	mu sync.Mutex // protects routeNotes
}

func newServer() *greetingServer {
	return &greetingServer{}
}

// GetFeature returns the feature at the given point.
func (s *greetingServer) Hello(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	msg := fmt.Sprintf("Hello, %s.", req.GetName())
	return &pb.GreetingResponse{Message: msg}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterGreetingServiceServer(s, newServer())

	log.Printf("Register reflection service")
	reflection.Register(s)

	log.Printf("gRPC Server start to listen on localhost:%d", *port)
	s.Serve(lis)
}
