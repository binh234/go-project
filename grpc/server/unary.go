package main

import (
	"context"

	pb "grpc-demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoPram) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
