package api

import (
	"errors"
	"net/http"

	"github.com/flightlogteam/flightlog-flights/src/logfile"
	"github.com/gorilla/mux"
	"github.com/klyngen/jsend"
)

func (a *FlightsAPI) mountLogfileRoutes(router *mux.Router) {
	router.HandleFunc("/logfile/igc", a.handleProcessIGCLogfile).Methods(http.MethodPost, http.MethodOptions)
}

func (a *FlightsAPI) handleProcessIGCLogfile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("logfile")

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	defer file.Close()

	var bytes []byte

	_, err = file.Read(bytes)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	reader, err := logfile.NewIGCLogfileReader(string(bytes))

	if err != nil {
		jsend.FormatResponse(w, errors.New("Unable to read the logfile"), jsend.BadRequest)
		return
	}
	records, err := reader.ReadRecords()

	if err != nil {
		jsend.FormatResponse(w, errors.New("Unable to read the logfile"), jsend.BadRequest)
		return
	}

	flightLog := logfile.NewFlightLog(records)

	jsend.FormatResponse(w, flightLog, jsend.Success)
}
