package main

import (
	"fmt"

	"github.com/JerryJeager/amor-rendezvous-backend/cmd"
	"github.com/JerryJeager/amor-rendezvous-backend/config"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
	fmt.Println("env and database initializaed successfully...")
	fmt.Println("starting to the amor-rendezvous server...")
}

func main() {
	cmd.ExecuteApiRoutes()
}
