package main

import (
	"context"
	"log"

	pb "../proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const PORT = "1234"

func main() {
	c, err := credentials.NewClientTLSFromFile("../conf/cert.pem", "localhost")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile Error: %v", err)
	}
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err %s", err)
	}
	defer conn.Close()

	client := pb.NewFindingServiceClient(conn)
	res, err := client.ReportMissing(context.Background(), &pb.FindingRequest{
		Name: "gRPC",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response is %s", res.GetMessage())
}
