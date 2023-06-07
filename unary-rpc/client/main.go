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
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	c := proto.NewMoneyTransactionClient(conn)

	transactionResponse, err := c.MakeTransaction(
		context.Background(),
		&proto.TransactionRequest{
			From:   "Mehul",
			To:     "Hiren",
			Amount: float32(120.15),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(transactionResponse.Confirmation)
}
