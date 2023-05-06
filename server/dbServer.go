package server

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"
)

func InitDatabase(config *viper.Viper) *sql.DB {
	connectionString := config.GetString(
		"database.connection_string")
	maxIdleConnection := config.GetInt(
		"database.max_idle_connections")
	maxOpenConnection := config.GetInt(
		"database.max_open_connections")
	connectionMaxLifeTime := config.GetDuration(
		"database.connection_max_lifetime")
	drvierName := config.GetString("database.driver_name")
	if connectionString == "" {
		log.Fatalf("Database connection string is missing")
	}

	dbHandler, err := sql.Open(drvierName, connectionString)
	if err != nil {
		log.Fatalf("Error while initilizing database: %v", err)
	}
	dbHandler.SetMaxIdleConns(maxIdleConnection)
	dbHandler.SetMaxOpenConns(maxOpenConnection)
	dbHandler.SetConnMaxLifetime(connectionMaxLifeTime)
	err = dbHandler.Ping()
	if err != nil {
		dbHandler.Close()
		log.Fatalf("Error while validating database: %v", err)
	}
	return dbHandler
}
