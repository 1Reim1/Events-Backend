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

	storage, err := storage.NewFileStorage(conf.DataFile)

	apiServer := apiserver.NewAPIServer(conf, storage)
	if apiServer.Start() != nil {
		log.Fatal(err)
	}
}
