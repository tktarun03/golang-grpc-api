package main

import (
	"context"
	"log"
	"net"

	pb "example.com/golang-grpc-api/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedItemServiceServer
}

func (s *server) GetItem(ctx context.Context, req *pb.ItemRequest) (*pb.ItemResponse, error) {
	return &pb.ItemResponse{Id: req.Id, Name: "Sample Item", Description: "This is a sample gRPC item"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterItemServiceServer(s, &server{})

	log.Println("gRPC Server is running on port 50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
