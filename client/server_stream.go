package main

import (
	"io"
	"log"

	pb "github.com/oigwebuike/go-grpc-init/proto"
	"golang.org/x/net/context"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}

		log.Println(message)

	}
	log.Printf("Streaming finished")

}
