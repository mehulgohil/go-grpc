package main

import (
	"fmt"
	"github.com/mehulgohil/go-grpc/bidirectional-streaming-rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewChatServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Fatalf("error opening stream: %v", err)
	}

	// Create a channel to receive messages from the server
	messageChannel := make(chan *proto.Message)

	// Start a goroutine to receive messages from the server
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("done")
				return
			}
			if err != nil {
				log.Fatalf("error receiving message: %v", err)
			}
			messageChannel <- message
		}
	}()

	// Send messages to the server
	for i := 0; i < 5; i++ {
		message := &proto.Message{
			UserId: strconv.Itoa(rand.Int()),
			Text:   "Hi",
		}
		if err := stream.Send(message); err != nil {
			log.Fatalf("error sending message: %v", err)
		}
		time.Sleep(time.Second) // Add some delay between sending messages
	}

	stream.CloseSend()

	for {
		message, ok := <-messageChannel
		if !ok {
			// Channel closed
			break
		}
		log.Printf("Received message: User ID=%s, Text=%s", message.UserId, message.Text)
	}
}
