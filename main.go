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

	s, err := storage.NewMySQLStorage(conf)
	if err != nil {
		log.Fatal(err)
	}

	apiServer := apiserver.NewAPIServer(conf, s)
	if err = apiServer.Start(); err != nil {
		log.Fatal(err)
	}
}
