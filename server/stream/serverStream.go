package main

import (
	"io"
	"log"
	"net"
	"time"

	pb "../../proto"
	"google.golang.org/grpc"
)

const STREAM_PORT = "7770"

type ContinuousStreamService struct{}

func main() {
	server := grpc.NewServer()
	pb.RegisterContinuousStreamServiceServer(server, &ContinuousStreamService{})

	listener, err := net.Listen("tcp", ":"+STREAM_PORT)
	if err != nil {
		log.Fatalf("listening error")
	}

	log.Printf("Start listening on %s", STREAM_PORT)

	server.Serve(listener)
}

func (s *ContinuousStreamService) ServerNotify(req *pb.ContinuousStreamRequest, stream pb.ContinuousStreamService_ServerNotifyServer) error {
	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.ContinuousStreamResponse{
			Chunk: &pb.ContinuousStreamChunk{
				Name: req.Chunk.Name,
				Id:   req.Chunk.Id + int32(n),
			},
		})

		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 1000)
	}
	return nil
}

func (s *ContinuousStreamService) ClientGiveIn(stream pb.ContinuousStreamService_ClientGiveInServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&pb.ContinuousStreamResponse{
					Chunk: &pb.ContinuousStreamChunk{
						Name: "gRPC CHUNK",
						Id:   2,
					}})
		}

		if err != nil {
			return err
		}

		log.Printf("Client Mode from Client: chunk.name: %s, chunk.Id: %d", req.Chunk.Name, req.Chunk.Id)
	}
}

func (s *ContinuousStreamService) HasConversation(stream pb.ContinuousStreamService_HasConversationServer) error {
	n := 0
	for n < 6 {
		err := stream.Send(&pb.ContinuousStreamResponse{
			Chunk: &pb.ContinuousStreamChunk{
				Name: "Conversation from server",
				Id:   int32(n),
			},
		})

		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("Conversation Mode from Client: name is %s and id is %d", req.Chunk.Name, req.Chunk.Id)
	}

	return nil
}
