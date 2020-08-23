package main

import (
	"context"
	"log"

	pb "../proto"
	"google.golang.org/grpc"
)

const PORT = "8882"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err %s", err)
	}
	defer conn.Close()

	client := pb.NewFindingServiceClient(conn)
	res, err := client.ReportMissing(context.Background(), &pb.FindingRequest{
		Name: "gRPC",
	})

	if err != nil {
		log.Fatalf("Client finding error")
	}
	log.Printf("response is %s", res.GetMessage())
}
