package main

import (
	"github.com/ak-karimzai/cp-db/cmd/config"
	"github.com/ak-karimzai/cp-db/cmd/server"
	"github.com/ak-karimzai/cp-db/internal/jwt"
	"github.com/ak-karimzai/cp-db/internal/logger"
	_ "github.com/lib/pq"
)

func main() {
	logger.Init()
	logger.GetLogger().Info("Starting uc app App")
	logger.GetLogger().Info("Initializing configuration")
	config := config.InitConfig("cp_db.toml")
	logger.GetLogger().Info("Initializing JWT secret key")
	jwt.InitJwtSerectKey(config)
	logger.GetLogger().Info("Initializing database server")
	dbHandler := server.InitDatabase(config)
	logger.GetLogger().Info("Initializing http server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
