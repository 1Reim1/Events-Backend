package main

import (
	"Events-Backend/apiserver"
	"Events-Backend/config"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	apiServer := apiserver.NewAPIServer(conf)
	if apiServer.Start() != nil {
		log.Fatal(err)
	}
}
