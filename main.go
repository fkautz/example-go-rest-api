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
	r.HandleFunc("/time", echoHandler).Methods("GET")
	r.HandleFunc("/ws/time", wsHandler).Methods("GET")
	http.Handle("/", r)
	err := http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		log.Println(err)
	}
}
