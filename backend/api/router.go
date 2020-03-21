package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// StartRouter will start the HTTP API server
func StartRouter(port, clientid, issuer string) error {
	router := mux.NewRouter().StrictSlash(true)

	o := oauth2provider{
		ClientID: clientid,
		Issuer:   issuer,
	}

	// definition of global middlewares to use
	router.Use(loggingMiddleware)
	router.Use(o.verificationMiddleware)

	// definition of routes
	router.HandleFunc("/", homeLink)
	return http.ListenAndServe(":"+port, router)
}
