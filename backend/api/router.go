package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// StartRouter will start the HTTP API server
func StartRouter(verifyToken bool, port, clientid, issuer string) error {
	router := mux.NewRouter().StrictSlash(true)

	// definition of global middlewares to use
	router.Use(loggingMiddleware)

	router.HandleFunc("/", homeLink)

	setupAPIRouter(router)

	// start the server and return the err in the end
	return http.ListenAndServe(":"+port, router)
}

func setupAPIRouter(r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()

	// user api
	api.HandleFunc("/users/{id}", getUser).Methods(http.MethodGet)
	api.HandleFunc("/users/{id}", updateUser).Methods(http.MethodPut)
	api.HandleFunc("/users/{id}", deleteUser).Methods(http.MethodDelete)
	api.HandleFunc("/user", createUser).Methods(http.MethodPost)

	//donations api
	api.HandleFunc("/donations/{id}", getDonation).Methods(http.MethodGet)
	api.HandleFunc("/donations/{id}", updateDonation).Methods(http.MethodPut)
	api.HandleFunc("/donations/{id}", deleteDonation).Methods(http.MethodDelete)
	api.HandleFunc("/donation", createDonation).Methods(http.MethodPost)
}
