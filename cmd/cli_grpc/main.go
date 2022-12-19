// Package main implements a client for Ping service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/aitrailblazer/ait-gcp-go-grpc/rpc/v1/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultMessage = "world"
)

var (
	addr    = flag.String("addr", "localhost:50051", "the address to connect to")
	message = flag.String("message", defaultMessage, "message to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAitrailblazerServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendPing(ctx, &pb.PingRequest{Message: *message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Pong: %s", r.GetPong())
}
