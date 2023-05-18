package main

import (
	"fmt"
	"github.com/mehulgohil/go-grpc/server-streaming-rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	c := proto.NewStreamServiceClient(conn)

	transactionResponse, err := c.GetUsers(
		context.Background(),
		&proto.Request{
			I: 1,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	for {
		resp, err := transactionResponse.Recv()
		if err == io.EOF {
			fmt.Println("done")
			return
		}
		if err != nil {
			log.Fatalf("can not receive %v", err)
		}
		log.Printf("Resp received: %s", resp.Username)
	}
}
