package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		POSTGRES_USER=env-pg-user;POSTGRES_PASSWORD=env-pg-password;REDIS_USER=env-user;REDIS_PASSWORD=env-redis-pass
	*/
	config, err := GetConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", config)
}
