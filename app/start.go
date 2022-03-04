package app

import (
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartPoint() {

	mux := mux2.NewRouter()
	mux.HandleFunc("/customers", fun).Methods(http.MethodGet)
	mux.HandleFunc("/customer/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
