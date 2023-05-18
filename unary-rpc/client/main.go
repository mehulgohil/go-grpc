package main

import (
	"context"
	"fmt"
	"github.com/mehulgohil/go-grpc/unary-rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// Set up connection with the grpc server
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	// Create a client instance
	c := proto.NewMoneyTransactionClient(conn)

	// Lets invoke the remote function from client on the server
	transactionResponse, err := c.MakeTransaction(
		context.Background(),
		&proto.TransactionRequest{
			From:   "mag",
			To:     "gam",
			Amount: float32(120.15),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(transactionResponse.Confirmation)
}
