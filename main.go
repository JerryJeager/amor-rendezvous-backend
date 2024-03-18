package main

import (
	"log"

	"github.com/JerryJeager/amor-rendezvous-backend/api"
	"github.com/JerryJeager/amor-rendezvous-backend/config"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
    log.Print("starting to the amor-rendezvous server...")

	api.ExecuteApiRoutes()
}