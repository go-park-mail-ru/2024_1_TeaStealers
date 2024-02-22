package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.HandleFunc("/ping", pingPongHeandler).Methods(http.MethodGet)
	log.Println("Start server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err.Error())
	}
}

func pingPongHeandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}
