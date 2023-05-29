package main

import (
	"context"

	pb "github.com/oigwebuike/go-grpc-init/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
