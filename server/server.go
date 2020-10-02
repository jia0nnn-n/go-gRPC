package main

import (
	"context"
	"log"
	"net"

	pb "../proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const PORT = "1234"

type FindingService struct{}

func (f *FindingService) ReportMissing(ctx context.Context, req *pb.FindingRequest) (*pb.FindingResponse, error) {
	return &pb.FindingResponse{Message: req.GetName() + "server"}, nil
}

func main() {
	c, err := credentials.NewServerTLSFromFile("../conf/cert.pem", "../conf/key.pem")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile Error: %v", err)
	}
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterFindingServiceServer(server, &FindingService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("listening error")
	}

	server.Serve(lis)
}
