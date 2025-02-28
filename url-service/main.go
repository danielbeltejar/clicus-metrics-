package main

import (
	"log"
	"net/http"
	"url-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	InitConfig()
	db := InitMongo()
	handlers.InitHandlers(db)

	r := mux.NewRouter()
	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/r", handlers.Redirect).Methods("GET")
	r.HandleFunc("/healthz", handlers.HealthCheck(db, nil)).Methods("GET")

	log.Println("URL service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
