package main

import (
	"context"
	"log"
	"time"

	pb "github.com/oigwebuike/go-grpc-init/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could nit send names: %v", err)

	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending request: %v", err)

		}

		log.Printf("Sent the request with name %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	log.Printf("Client sending finished")

	if err != nil {
		log.Fatalf("Error while recieving stream: %v", err)
	}

	log.Printf("%v", res.Messages)

}
