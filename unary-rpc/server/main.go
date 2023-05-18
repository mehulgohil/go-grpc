package main

import (
	"fmt"
	"github.com/mehulgohil/go-grpc/unary-rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedMoneyTransactionServer
}

func (s server) MakeTransaction(ctx context.Context, in *proto.TransactionRequest) (*proto.TransactionResponse, error) {
	fmt.Println("Got amount ", in.Amount)
	fmt.Println("Got from ", in.From)
	fmt.Println("For ", in.To)
	return &proto.TransactionResponse{Confirmation: true}, nil
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	proto.RegisterMoneyTransactionServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
