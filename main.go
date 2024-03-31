package main

import (
	"log"

	"github.com/JerryJeager/amor-rendezvous-backend/cmd"
	"github.com/JerryJeager/amor-rendezvous-backend/config"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
	log.Print("env and database initializaed successfully...")
}

func main() {
    log.Print("starting to the amor-rendezvous server...")

	cmd.ExecuteApiRoutes()
}