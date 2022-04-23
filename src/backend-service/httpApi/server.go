package httpApi

import (
	"context"
	"log"
	"net"

	"github.com/Todo-Task-Management/server/helloworld_pb"
	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, req *helloworld_pb.HelloWorldRequest) (*helloworld_pb.HelloWorldResponse, error) {
	return helloworld_pb.HelloWorldResponse{Message: "hello " + req.Name}, nil
}

func main() {
	grpcServer := grpc.NewServer()

	listen, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:5000 %v", err)
	}
	server := Server{}
	helloworld_pb.RegisterHelloWorldServiceServer(grpcServer, &server)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
