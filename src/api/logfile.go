package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

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

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, file); err != nil {
		jsend.FormatResponse(w, "Unable to read the file", jsend.BadRequest)
		return
	}

	fmt.Println(buf)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	flightLog, err := a.logfileService.ProcessIGCLogfile(string(buf.Bytes()))

	if err != nil {
		jsend.FormatResponse(w, errors.New("Unable to read the logfile"), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, flightLog, jsend.Success)
}
