package main

import (
	"log"

	"github.com/ghivarra/app/module/router"
	"github.com/ghivarra/app/server"
	"github.com/joho/godotenv"
)

func main() {

	// load environment
	errEnv := godotenv.Load("./.env")
	if errEnv != nil {
		log.Fatalf("Cannot load dotenv/.env file. Reason: %v", errEnv)
	}

	// load router
	router := router.RouteRegister()

	// run server
	server.Run(router)
}
