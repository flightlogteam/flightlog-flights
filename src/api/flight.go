package api

import (
	"encoding/json"
	"net/http"

	"github.com/flightlogteam/flightlog-flights/src/flight"
	"github.com/gorilla/mux"
	"github.com/klyngen/jsend"
)

func (a *FlightsAPI) mountFlightRoutes(router *mux.Router) {
	router.HandleFunc("/flight/{id}", a.handleGetFlightByID).Methods("GET")
	router.HandleFunc("/flight", a.handleCreateFlight).Methods("POST", http.MethodOptions)
	router.HandleFunc("/flight", a.handleUpdateFlight).Methods("PUT", http.MethodOptions)
}

func (a *FlightsAPI) handleGetFlightByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	flight, err := a.flightService.FlightById(id)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, flight, jsend.Success)
}

func (a *FlightsAPI) handleCreateFlight(w http.ResponseWriter, r *http.Request) {
	var flight flight.Flight

	if json.NewDecoder(r.Body).Decode(&flight) != nil {
		jsend.FormatResponse(w, "Unreadable JSON", jsend.BadRequest)
		return
	}

	id, err := a.flightService.Create(&flight)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, id, jsend.Success)
}

func (a *FlightsAPI) handleUpdateFlight(w http.ResponseWriter, r *http.Request) {
	var flight flight.Flight

	if json.NewDecoder(r.Body).Decode(&flight) != nil {
		jsend.FormatResponse(w, "Unreadable JSON", jsend.BadRequest)
		return
	}

	err := a.flightService.Update(&flight)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, nil, jsend.Success)
}
