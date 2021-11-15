package main

import (
	"context"
	"fmt"
	pb "grpc_stream_example/proto"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type MyServer struct {
	Name string
	pb.UnimplementedIncrementorServer
}

func (s *MyServer) Inc(ctx context.Context, req *pb.NumArgument) (*pb.NumResult, error) {
	log.Printf("Inc method called")

	return &pb.NumResult{
		N:           req.GetN() + 1,
		ServiceName: s.Name,
	}, nil
}

func (s *MyServer) NInc(req *pb.NNumArgument, stream pb.Incrementor_NIncServer) error {
	log.Printf("NInc method called")

	// go func(stream pb.Incrementor_NIncServer) {
	count := int(req.GetTimes())

	result := &pb.NumResult{
		N:           req.GetN(),
		ServiceName: s.Name,
	}
	ctx := stream.Context()
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			log.Printf("Context canceled: %v", ctx.Err())
			return ctx.Err()
		default:
			// do nothing
			break
		}
		result.N++
		if err := stream.Send(result); err != nil {
			log.Printf("Response sending failed. Error: %v", err)
			return err
		}
		log.Println("Response successfully sent on stream")
		time.Sleep(time.Second)
	}
	// }(stream)

	return nil
}

func (s *MyServer) Sum(stream pb.Incrementor_SumServer) error {
	log.Printf("Sum method called")

	result := &pb.NumResult{
		ServiceName: s.Name,
	}

	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Context canceled: %v", ctx.Err())
			return fmt.Errorf("Sum context error")
		default:
			// do nothing
			break
		}
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Recv error: %v", err)
			return err
		}

		log.Printf("Request received: %+v", req)
		result.N += req.GetN()
	}

	// Need send response
	err := stream.SendAndClose(result)
	if err != nil {
		log.Printf("Response send error: %v", err)
		return err
	}
	return nil
}

func (s *MyServer) StreamSum(stream pb.Incrementor_StreamSumServer) error {
	log.Printf("StreamSum method called")

	result := &pb.NumResult{
		ServiceName: s.Name,
	}

	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Context canceled: %v", ctx.Err())
			return fmt.Errorf("StreamSum context error")
		default:
			// do nothing
			break
		}
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Recv error: %v", err)
			return err
		}

		log.Printf("Request received: %+v", req)
		result.N += req.GetN()

		if err := stream.Send(result); err != nil {
			log.Printf("Send error: %v", err)
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Too few arguments!")
	}

	port := os.Args[1]
	name := os.Args[2]

	rand.Seed(time.Now().Unix())

	listener, err := net.Listen("tcp", port)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	// Create server instance
	server := &MyServer{
		Name: name,
	}

	pb.RegisterIncrementorServer(grpcServer, server)

	log.Printf("Start gRPC service...")
	log.Fatal(grpcServer.Serve(listener))
}
