package main

import (
	"log"

	"github.com/ak-karimzai/cp-db/config"
	"github.com/ak-karimzai/cp-db/jwt"
	"github.com/ak-karimzai/cp-db/server"
	_ "github.com/lib/pq"
)

func main() {
	log.Printf("Starting runners App")
	log.Printf("Initializing configuration")
	config := config.InitConfig("cp_db.toml")
	log.Printf("Initializing JWT secret key")
	jwt.InitJwtSerectKey(config)
	log.Printf("Initializing database server")
	dbHandler := server.InitDatabase(config)
	log.Printf("Initializing http server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
