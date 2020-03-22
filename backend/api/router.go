package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// StartRouter will start the HTTP API server
func StartRouter(port string) error {
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
	api.HandleFunc("/users", listUsers).Methods(http.MethodGet)
	api.HandleFunc("/user", createUser).Methods(http.MethodPost)

	//donations api
	api.HandleFunc("/donations/{id}", getDonation).Methods(http.MethodGet)
	api.HandleFunc("/donations/{id}", updateDonation).Methods(http.MethodPut)
	api.HandleFunc("/donations/{id}", deleteDonation).Methods(http.MethodDelete)
	api.HandleFunc("/donations", listDonations).Methods(http.MethodGet)
	api.HandleFunc("/donations/byReceiver/{id}", listDonationsByReceiver).Methods(http.MethodGet) // this {id} is meant to be an existing User ID
	api.HandleFunc("/donation", createDonation).Methods(http.MethodPost)

	// company api
	api.HandleFunc("/companies/{id}", getCompany).Methods(http.MethodGet)
	api.HandleFunc("/companies/{id}", updateCompany).Methods(http.MethodPut)
	api.HandleFunc("/companies/{id}", deleteCompany).Methods(http.MethodDelete)
	api.HandleFunc("/companies", listCompanies).Methods(http.MethodGet)
	api.HandleFunc("/company", createCompany).Methods(http.MethodPost)

	// image api
	api.HandleFunc("/images/{id}", getImage).Methods(http.MethodGet)
	api.HandleFunc("/images/{id}", updateImage).Methods(http.MethodPut)
	api.HandleFunc("/images/{id}", deleteImage).Methods(http.MethodDelete)
	api.HandleFunc("/images", listImages).Methods(http.MethodGet)
	api.HandleFunc("/image", createImage).Methods(http.MethodPost)
}
