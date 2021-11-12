package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"
	pb "twirp_example/proto"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type MyServer struct {
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

	mux := http.NewServeMux()

	// Create server instance
	server := &MyServer{}

	twirpHandler := pb.NewMySuperServiceServer(server)

	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8080", mux)
}
