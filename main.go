package main

import (
	"log"
	"net/http"

	"github.com/kwandapchumba/changelog/router"
	"github.com/kwandapchumba/changelog/utils"
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

	log.Printf("Server running on http://localhost%s", config.ConnectionPort)

	log.Fatal(server.ListenAndServe())
}
