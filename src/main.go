package main

import (
	"log"
	"os"

	"github.com/flightlogteam/flightlog-flights/src/flight"
)

func main() {
	port := ":61226"
	config := getConfiguration()

	if !config.IsValidConfiguration() {
		log.Fatalf("No valid database configuration present \n")
	}

	repo, err := flight.NewRepository(config.username, config.password, config.hostname, config.port, config.database)

	if err != nil {
		log.Fatalf("cannot create connection to database: %v, on hostname %v, with user %v. Config gave following error:\n %v \n", config.database, config.hostname, config.username, err)
		return
	}

}

func getConfiguration() databaseConfiguration {
	return databaseConfiguration{
		password: os.Getenv("DATABASE_PASSWORD"),
		username: os.Getenv("DATABASE_USERNAME"),
		port:     os.Getenv("DATABASE_PORT"),
		hostname: os.Getenv("DATABASE_HOSTNAME"),
		database: os.Getenv("DATABASE_DATABASE"),
	}
}

type databaseConfiguration struct {
	password string
	username string
	port     string
	hostname string
	database string
}

func (c *databaseConfiguration) IsValidConfiguration() bool {
	if len(c.password) > 0 && len(c.username) > 0 {
		return true
	}
	return false
}
