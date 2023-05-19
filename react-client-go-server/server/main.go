package main

import (
	"fmt"
	"github.com/mehulgohil/go-grpc/unary-rpc/react-client-go-server/server/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type calculatorServer struct {
	proto.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	result := req.Operand1 + req.Operand2
	fmt.Println("result add", result)
	return &proto.AddResponse{Result: result}, nil
}

func (s *calculatorServer) Subtract(ctx context.Context, req *proto.SubtractRequest) (*proto.SubtractResponse, error) {
	result := req.Operand1 - req.Operand2
	fmt.Println("result sub", result)
	return &proto.SubtractResponse{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterCalculatorServer(grpcServer, &calculatorServer{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
