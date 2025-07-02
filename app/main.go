package main

import (
	"log"
	"path"

	"github.com/ghivarra/app/module/router"
	"github.com/ghivarra/app/server"
	"github.com/joho/godotenv"
)

var ROOT string

func main() {

	// put root
	ROOT = path.Clean("/")

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
