package main

import (
	"github.com/mehulgohil/go-grpc/bidirectional-streaming-rpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// ChatServer implements the ChatServiceServer interface
type ChatServer struct {
	proto.UnimplementedChatServiceServer
}

// SendMessage implements the SendMessage gRPC method
func (s *ChatServer) SendMessage(stream proto.ChatService_SendMessageServer) error {
	for {
		message, err := stream.Recv()
		if err != nil {
			return err
		}

		reply := &proto.Message{
			UserId: message.UserId,
			Text:   "Bye from server",
		}

		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	chatServer := &ChatServer{}
	proto.RegisterChatServiceServer(grpcServer, chatServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
