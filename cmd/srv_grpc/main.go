// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/aitrailblazer/ait-gcp-go-grpc/rpc/v1/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const VERSION = "1.02"

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedAitrailblazerServiceServer
}

// SendPing
func (s *server) SendPing(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	var i int32 = 369
	msec := timestamppb.Now()

	var v string = VERSION
	log.Printf("AitrailblazerServiceSend VERSION %s ", VERSION) // <3>
	message := in.GetMessage()
	log.Printf("Received message: %v", message)

	pong := pb.Pong{
		Index:      i,
		Message:    message,
		ReceivedOn: msec,
		Ver:        v,
	}
	pingResponse := pb.PingResponse{Pong: &pong}
	return &pingResponse, nil
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAitrailblazerServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
