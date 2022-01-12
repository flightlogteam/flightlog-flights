package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flightlogteam/flightlog-flights/src/flight"
	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/flightlogteam/flightlog-flights/src/start"
	"github.com/gorilla/mux"
	"github.com/klyngen/jsend"
)

// FlightsAPI defines dependencies for our API
type FlightsAPI struct {
	router          *mux.Router
	locationService location.Service
	flightService   flight.Service
	startService    start.Service
}

// NewFlightsApi creates an api with a functioning router
func NewFlightsAPI(locationService location.Service, flightService flight.Service, startService start.Service) FlightsAPI {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsend.FormatResponse(w, fmt.Sprintf("Flight-API does not recognize the path: %v", r.URL.Path), jsend.NotFound)
	})

	api := FlightsAPI{
		locationService: locationService,
		flightService:   flightService,
		startService:    startService,
		router:          router,
	}
	api.mountFlightRoutes(router)
	api.mountStartRoutes(router)
	api.mountLocationRoutes(router)

	return api
}

func (f *FlightsAPI) StartAPI() {
	log.Printf("Starting Flights-API on port %v", "61227")
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", "61227"), f.router)

	if err != nil {
		log.Fatalf("Unable to start the all important API: %v \n", err)
	}
}
