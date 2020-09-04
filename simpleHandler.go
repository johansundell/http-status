package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func init() {
	routes = append(routes, Route{"simpleHandler", "GET", "/{argument}", simpleHandler})
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if settings.Debug {
		logger.Info(vars)
	}
	code, err := strconv.Atoi(vars["argument"])
	if err != nil {
		http.Error(w, "Not found status", 500)
		return
	}

	http.Error(w, http.StatusText(code), code)
}
