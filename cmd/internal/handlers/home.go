package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home page")
	return
}
