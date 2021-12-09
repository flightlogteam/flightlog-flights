package main

import (
	"log"
	"os"

	"github.com/flightlogteam/flightlog-flights/src/api"
	"github.com/flightlogteam/flightlog-flights/src/common"
	"github.com/flightlogteam/flightlog-flights/src/flight"
	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/flightlogteam/flightlog-flights/src/start"
)

func main() {
	log.Println("Starting flight-service")
	//port := ":61227"
	config, serviceConfig := getConfiguration()

	if !config.IsValidConfiguration() {
		log.Fatalf("No valid database configuration present \n")
	}

	common.EnsureDatabaseExists(config.username, config.password, config.hostname, config.port, config.database)

	connection, err := common.CreateGormContext(config.username, config.password, config.hostname, config.port, config.database)

	if err != nil {
		log.Fatalf("cannot create connection to database: %v, on hostname %v, with user %v. Config gave following error:\n %v \n", config.database, config.hostname, config.username, err)
		return
	}

	locationRepo, err := location.NewRepository(connection)
	flightRepo, err := flight.NewRepository(connection)
	startRepo, err := start.NewRepository(connection)

	if err != nil {
		log.Fatalf("Unable to auto migrate database %v", err)
		return
	}

	locationService := location.NewService(locationRepo)
	flightService, err := flight.NewService(serviceConfig.UserserviceUrl, serviceConfig.UserservicePort, flightRepo, serviceConfig, locationService)
	startService := start.NewService(startRepo, locationService)

	if err != nil {
		log.Printf("Unable instantiate flightservice. %v\n", err)
	}

	flightAPI := api.NewFlightsAPI(locationService, flightService, startService)
	flightAPI.StartAPI()
}

func getConfiguration() (databaseConfiguration, common.ServiceConfig) {
	return databaseConfiguration{
			password: os.Getenv("DATABASE_PASSWORD"),
			username: os.Getenv("DATABASE_USERNAME"),
			port:     os.Getenv("DATABASE_PORT"),
			hostname: os.Getenv("DATABASE_HOSTNAME"),
			database: os.Getenv("DATABASE_DATABASE"),
		}, common.ServiceConfig{
			UserserviceUrl:  os.Getenv("SERVICE_USERSERVICE_URL"),
			UserservicePort: os.Getenv("SERVICE_USERSERVICE_PORT"),
			Mode:            os.Getenv("SERVICE_MODE"),
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
