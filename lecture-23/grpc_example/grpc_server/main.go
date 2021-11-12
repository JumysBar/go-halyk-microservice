package main

import (
	"context"
	pb "grpc_example/proto"
	"log"
	"math/rand"
	"net"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type MyServer struct {
	pb.UnimplementedMySuperServiceServer
}

func (s *MyServer) AddUser(ctx context.Context, user *pb.User) (*pb.Status, error) {
	log.Printf("Method for adding user. User info: %+v", user)
	return &pb.Status{
		Code:        0,
		Description: "SUCCESS",
	}, nil
}

func (s *MyServer) GetRandomNumber(ctx context.Context, e *empty.Empty) (*pb.NumResult, error) {
	log.Printf("Getting random number RPC")
	return &pb.NumResult{
		Result: int64(rand.Int()),
	}, nil
}

func main() {
	rand.Seed(time.Now().Unix())

	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	// Create server instance
	server := &MyServer{}

	pb.RegisterMySuperServiceServer(grpcServer, server)

	log.Printf("Start gRPC service...")
	log.Fatal(grpcServer.Serve(listener))
}
