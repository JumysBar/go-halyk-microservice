package main

import (
	"context"
	pb "grpc_example/proto"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
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

	client := pb.NewMySuperServiceClient(conn)

	user := &pb.User{
		Id:        1,
		FirstName: "First",
		LastName:  "Second",
		Tags:      []string{"user", "example", "pornhub"},
		Color:     pb.Color_Yellow,
		Favorite: &pb.Film{
			Name: "Индиана Джонс: В поисках утраченного ковчега",
			Year: 1981,
		},
		Credentials: &pb.User_Credentials{
			Login:    "admin",
			Password: "admin",
		},

		Error: &pb.User_Description{
			Description: "Error!",
		},
	}

	response, err := client.AddUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("AddUser response: %+v", response)

	numResponse, err := client.GetRandomNumber(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GetRandomNumber response: %+v", numResponse)
}
