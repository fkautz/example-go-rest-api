package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Running")
	r := mux.NewRouter()
	r.HandleFunc("/", structuredHandler).Methods("GET")
	r.HandleFunc("/hello/{person}", helloHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:8080", r)
}
