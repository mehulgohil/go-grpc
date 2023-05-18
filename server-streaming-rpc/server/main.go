package main

import (
	"fmt"
	"github.com/mehulgohil/go-grpc/server-streaming-rpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServer struct {
	proto.UnimplementedStreamServiceServer
}

func (s *UserServer) GetUsers(req *proto.Request, stream proto.StreamService_GetUsersServer) error {
	fmt.Println("request", req.GetI())

	users := []proto.Response{
		{
			Username: "mag",
		},
		{
			Username: "gohil",
		},
	}

	for _, user := range users {
		if err := stream.Send(&user); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the UserServer with the gRPC server
	userServer := &UserServer{}
	proto.RegisterStreamServiceServer(grpcServer, userServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
