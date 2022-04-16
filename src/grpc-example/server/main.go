package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Todo-Task-Management/src/grpc-example/users_pb"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) ScoreAvg(ctx context.Context, userRequest *users_pb.UserRequest) (*users_pb.UserResponse, error) {
	name, height := userRequest.GetName(), userRequest.GetHeight()
	eyeColour, hasFacebook := userRequest.GetEyeColour(), userRequest.GetHasFacebook()
	tests := userRequest.GetTest()
	var testsTotalScore int32 = 0

	testNames := ""

	for _, test := range tests {
		testsTotalScore += test.Score
		testNames += fmt.Sprintf("%s,", test.Name)
	}
	// cut the last comma
	testNames = testNames[:len(testNames)-1]

	message := fmt.Sprintf("Name: %s, Height: %f, EyeColour: %s, Has Facebook: %v, Test Names: %s", name, height, eyeColour, hasFacebook, testNames)

	response := &users_pb.UserResponse{
		ScoreAverage: testsTotalScore / int32(len(tests)),
		UserMessage:  message,
	}

	return response, nil

}

func main() {
	grpcServer := grpc.NewServer()

	listen, err := net.Listen("tcp", "0.0.0.0:3001")
	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:3001 %v", err)
	}

	fmt.Println("worked!")

	server := Server{}

	users_pb.RegisterUserServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

