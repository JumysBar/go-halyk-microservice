package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {

	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n\n", config)

	fmt.Printf("%#v\n", viper.Get("postgres"))
	fmt.Printf("%#v\n", viper.Get("some_port"))
	fmt.Printf("%#v\n", viper.Get("user"))
	fmt.Printf("%#v\n", viper.Get("password"))

	fmt.Printf("%#v\n", viper.AllSettings())
}
