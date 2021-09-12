package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/gorilla/mux"
	"github.com/klyngen/jsend"
)

func (a *FlightsAPI) mountLocationRoutes(router *mux.Router) {
	router.HandleFunc("/location", a.handleGetLocation).Methods("GET")
	router.HandleFunc("/location", a.handleCreateLocation).Methods("POST")
	router.HandleFunc("/location", a.handleUpdateLocation).Methods("PUT")
	router.HandleFunc("/location/{id}", a.handleGetLocationByID).Methods("GET")
}

func (a *FlightsAPI) handleGetLocation(w http.ResponseWriter, r *http.Request) {
	locations, err := a.locationService.GetAllLocations()

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.InternalServerError)
		return
	}

	jsend.FormatResponse(w, locations, jsend.Success)
}

func (a *FlightsAPI) handleGetLocationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)

	if err != nil {
		jsend.FormatResponse(w, "ID is not and integer", jsend.BadRequest)
		return
	}

	location, err := a.locationService.GetLocationById(uint(id))

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, location, jsend.Success)
}

func (a *FlightsAPI) handleCreateLocation(w http.ResponseWriter, r *http.Request) {
	var location location.Location

	if json.NewDecoder(r.Body).Decode(&location) != nil {
		jsend.FormatResponse(w, "Unreadable JSON", jsend.BadRequest)
		return
	}
	id, err := a.locationService.CreateLocation(&location)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, id, jsend.Success)
}

func (a *FlightsAPI) handleUpdateLocation(w http.ResponseWriter, r *http.Request) {
	var location location.Location

	if json.NewDecoder(r.Body).Decode(&location) != nil {
		jsend.FormatResponse(w, "Unreadable JSON", jsend.BadRequest)
		return
	}
	err := a.locationService.UpdateLocation(&location)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, nil, jsend.Success)
}
