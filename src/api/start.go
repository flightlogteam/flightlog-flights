package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flightlogteam/flightlog-flights/src/start"
	"github.com/gorilla/mux"
	"github.com/klyngen/jsend"
)

func (a *FlightsAPI) mountStartApi(router *mux.Router) {
	router.HandleFunc("/start", a.handleStartCreation).Methods("POST")
	router.HandleFunc("/start", a.handleStartUpdate).Methods("PUT")
	router.HandleFunc("/start/{id}", a.handleStartDeletion).Methods("DELETE")
}

func (a *FlightsAPI) handleStartCreation(w http.ResponseWriter, r *http.Request) {
	var start start.Start

	if json.NewDecoder(r.Body).Decode(&start) != nil {
		jsend.FormatResponse(w, "Unreadable JSON", jsend.BadRequest)
		return
	}

	id, err := a.startService.Create(&start)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, id, jsend.Success)
}

func (a *FlightsAPI) handleStartUpdate(w http.ResponseWriter, r *http.Request) {
	var start start.Start

	if json.NewDecoder(r.Body).Decode(&start) != nil {
		jsend.FormatResponse(w, "Unreadable JSON", jsend.BadRequest)
		return
	}

	err := a.startService.Update(&start)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, true, jsend.Success)
}

func (a *FlightsAPI) handleStartDeletion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)

	if err != nil {
		jsend.FormatResponse(w, "ID is not and integer", jsend.BadRequest)
		return
	}

	err = a.startService.Delete(uint(id))

	if err != nil {
		jsend.FormatResponse(w, "Unable to delete the given start", jsend.InternalServerError)
		return
	}

	jsend.FormatResponse(w, true, jsend.Success)
}
