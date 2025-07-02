package main

import (
	"log"
	"os"

	"github.com/ghivarra/app/common"
	"github.com/ghivarra/app/router"
	"github.com/ghivarra/app/server"
	"github.com/joho/godotenv"
)

func main() {

	// put root
	common.ROOT, _ = os.Getwd()

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
