package main

import (
	"log"
	"net/http"

	"github.com/JulyInSummer/forms_clone.git/cmd/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Home)

	log.Println("Starting server")
	http.ListenAndServe(":8080", router)
}
