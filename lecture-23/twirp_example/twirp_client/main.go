package main

import (
	"context"
	"log"
	"net/http"
	pb "twirp_example/proto"

	"github.com/golang/protobuf/ptypes/empty"
)

func main() {
	client := pb.NewMySuperServiceProtobufClient("http://localhost:8080", &http.Client{})

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
