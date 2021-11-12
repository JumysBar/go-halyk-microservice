package main

import (
	"log"
	pb "proto_example/proto"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
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

	someProtoStruct := &pb.Film{
		Name: "Зеленый слоник",
		Year: 1999,
	}

	someAny := &anypb.Any{}
	someAny.MarshalFrom(someProtoStruct)

	user.Dynamic = someAny

	data, err := proto.Marshal(user)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Marshalled data len: %d", len(data))

	unmarshalledUser := &pb.User{}

	if err := proto.Unmarshal(data, unmarshalledUser); err != nil {
		log.Fatal(err)
	}

	log.Printf("Unmarshalled data: %+v", unmarshalledUser)
}
