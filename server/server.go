package main

import (
	"context"
	"log"
	"net"

	pb "../proto"
	"google.golang.org/grpc"
)

type FindingService struct{}

func (s *FindingService) Finding(ctx context.Context, r *pb.FindingRequest) (*pb.FindingResponse, error) {
	return &pb.FindingResponse{Message: "Server"}, nil
}

const PORT = "8882"

func main() {
	server := grpc.NewServer()

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("listening error")
	}

	server.Serve(lis)
}
