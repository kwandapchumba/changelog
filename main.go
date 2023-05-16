package main

import (
	"log"
	"net/http"

	"github.com/kwandapchumba/prioritize/router"
	"github.com/kwandapchumba/prioritize/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	server := &http.Server{
		Addr:    config.ConnectionPort,
		Handler: router.Router(),
	}

	log.Fatal(server.ListenAndServe())
}
