package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct for the service
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes for the servcie web handlers
type Routes []Route

// NewRouter creates a new web handler
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = wwwLogger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func wwwLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if settings.Debug {
			logger.Info(name + " " + r.RequestURI + " " + r.RemoteAddr + " " + r.Method)
		}
		w.Header().Set("X-Version", appVersionStr)
		inner.ServeHTTP(w, r)
	})
}
