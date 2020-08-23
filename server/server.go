package main

import (
	"context"
	"log"
	"net"

	pb "../proto"
	"google.golang.org/grpc"
)

const PORT = "1234"

type FindingService struct{}

func (f *FindingService) ReportMissing(ctx context.Context, req *pb.FindingRequest) (*pb.FindingResponse, error) {
	return &pb.FindingResponse{Message: req.GetName() + "server"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterFindingServiceServer(server, &FindingService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("listening error")
	}

	server.Serve(lis)
}
