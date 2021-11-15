package main

import (
	"context"
	pb "grpc_stream_example/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	log.Println("gRPC connect was created")

	client := pb.NewIncrementorClient(conn)

	// Inc RPC

	for i := 0; i < 3; i++ {
		response, err := client.Inc(context.Background(), &pb.NumArgument{
			N: 10,
		})

		if err != nil {
			log.Fatalf("Inc error: %v", err)
		}

		log.Printf("Inc response: %+v", response)
	}

	// NInc RPC

	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	// defer cancel()
	NIncStream, err := client.NInc(ctx, &pb.NNumArgument{
		N:     10,
		Times: 10,
	})

	if err != nil {
		log.Fatalf("NInc error: %v", err)
	}

	for {
		resp, err := NIncStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Recv error: %v", err)
		}

		log.Printf("NInc response: %+v", resp)
	}

	log.Println("End of NInc RPC")

	// Sum RPC

	// ctx = context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	// defer cancel()
	SumStream, err := client.Sum(ctx)
	if err != nil {
		log.Fatalf("Sum error: %v", err)
	}

	for i := 0; i < 10; i++ {
		if err := SumStream.Send(&pb.NumArgument{
			N: 10,
		}); err != nil {
			log.Printf("Send error: %v", err)
			err := SumStream.RecvMsg(nil)
			if err != nil {
				log.Fatalf("Stream error: %v", err)
			}
		}

		log.Printf("Message successfully sent")
		time.Sleep(time.Second)
	}

	resp, err := SumStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Close and receive error: %v", err)
	}

	log.Printf("Sum response: %+v", resp)

	// StreamSum RPC

	// ctx = context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	// defer cancel()
	StreamSumStream, err := client.StreamSum(ctx)
	if err != nil {
		log.Fatalf("StreamSum error: %v", err)
	}

	for i := 0; i < 10; i++ {
		if err := StreamSumStream.Send(&pb.NumArgument{
			N: 10,
		}); err != nil {
			log.Printf("Send error: %v", err)
			err := StreamSumStream.RecvMsg(nil)
			if err != nil {
				log.Fatalf("Stream error: %v", err)
			}
		}

		log.Printf("Message successfully sent")

		resp, err := StreamSumStream.Recv()
		if err != nil {
			log.Fatalf("Recv error: %v", err)
		}

		log.Printf("Message successfully received: %+v", resp)

		time.Sleep(time.Second)
	}

	log.Println("All RPC called")
}
