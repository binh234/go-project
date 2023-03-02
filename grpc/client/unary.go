package main

import (
	"context"
	"log"
	"time"

	pb "grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoPram{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(res.Message)
}
