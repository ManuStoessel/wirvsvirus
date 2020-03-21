package api

import (
	"fmt"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the wirvsvirus backend!")
}
