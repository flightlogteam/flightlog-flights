package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flightlogteam/flightlog-flights/src/flight"
	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/flightlogteam/flightlog-flights/src/logfile"
	"github.com/flightlogteam/flightlog-flights/src/start"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/klyngen/jsend"
)

// FlightsAPI defines dependencies for our API
type FlightsAPI struct {
	router          *mux.Router
	locationService location.Service
	flightService   flight.Service
	startService    start.Service
	logfileService  logfile.Service
}

// NewFlightsApi creates an api with a functioning router
func NewFlightsAPI(locationService location.Service, flightService flight.Service, startService start.Service, logservice logfile.Service) FlightsAPI {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsend.FormatResponse(w, fmt.Sprintf("Flight-API does not recognize the path: %v", r.URL.Path), jsend.NotFound)
	})

	api := FlightsAPI{
		locationService: locationService,
		flightService:   flightService,
		startService:    startService,
		router:          router,
		logfileService:  logservice,
	}
	api.mountFlightRoutes(router)
	api.mountStartRoutes(router)
	api.mountLocationRoutes(router)
	api.mountLogfileRoutes(router)

	router.Use(mux.CORSMethodMiddleware(router))

	return api
}

func (f *FlightsAPI) StartAPI() {
	log.Printf("Starting Flights-API on port %v", "61227")
	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", "61227"), handlers.CORS(headersOK, originsOK, methodsOK)(f.router))

	if err != nil {
		log.Fatalf("Unable to start the all important API: %v \n", err)
	}
}
