package main

import (
	"Events-Backend/apiserver"
	"Events-Backend/config"
	"Events-Backend/storage"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	s, err := storage.NewFileStorage(conf.DataFile)
	if err != nil {
		log.Fatal(err)
	}

	apiServer := apiserver.NewAPIServer(conf, s)
	if apiServer.Start() != nil {
		log.Fatal(err)
	}
}
