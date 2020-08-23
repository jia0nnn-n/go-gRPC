package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	pb "../../proto"
	"google.golang.org/grpc"
)

const PORT = "7770"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err %s", err)
	}
	defer conn.Close()

	client := pb.NewContinuousStreamServiceClient(conn)
	rand.Seed(time.Now().Unix())

	err = printNotify(client, &pb.ContinuousStreamRequest{Chunk: &pb.ContinuousStreamChunk{Name: "Notify", Id: rand.Int31n(100)}})
	if err != nil {
		logErrMessage("printNotify", err.Error())
	}

	err = printGiveIn(client, &pb.ContinuousStreamRequest{Chunk: &pb.ContinuousStreamChunk{Name: "GiveIn", Id: 23}})
	if err != nil {
		logErrMessage("printGiveIn", err.Error())
	}

	err = printConversation(client, &pb.ContinuousStreamRequest{Chunk: &pb.ContinuousStreamChunk{Name: "Conversation", Id: 23}})
	if err != nil {
		logErrMessage("printConversation", err.Error())
	}
}

func printNotify(client pb.ContinuousStreamServiceClient, req *pb.ContinuousStreamRequest) error {
	stream, err := client.ServerNotify(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		log.Printf("Server Mode from Server: name is %s and id is %d", res.Chunk.Name, res.Chunk.Id)
	}

	return nil
}

func printGiveIn(client pb.ContinuousStreamServiceClient, r *pb.ContinuousStreamRequest) error {
	stream, err := client.ClientGiveIn(context.Background())

	if err != nil {
		return err
	}

	for n := 0; n < 5; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 1000)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("Client Mode from Server One Time: name is %s and id is %d", res.Chunk.Name, res.Chunk.Id)
	return nil
}

func printConversation(client pb.ContinuousStreamServiceClient, r *pb.ContinuousStreamRequest) error {
	stream, err := client.HasConversation(context.Background())
	if err != nil {
		return err
	}

	n := 0
	for {
		err = stream.Send(r)
		if err != nil {
			return err
		}

		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		n++

		log.Printf("Conversation Mode from Server: name is %s and id is %d", res.Chunk.Name, res.Chunk.Id)
		time.Sleep(time.Millisecond * 1000)
	}
	stream.CloseSend()

	return nil
}

func logErrMessage(funcName string, errorMessage string) {
	var b bytes.Buffer
	b.WriteString("FAILED --- ")
	b.WriteString(funcName + "\n")
	b.WriteString(errorMessage)
	log.Fatal(b.String())
}
