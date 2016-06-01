package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m := make(map[string]string)
	m["hello"] = vars["person"]
	js, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.Write(js)
}

func structuredHandler(w http.ResponseWriter, r *http.Request) {
	sampleResponse := SampleResponse{
		Hello: "world",
		Count: 10,
		Address: FullAddress{
			StreetNumber: 1234,
			Street:       "Easy St.",
			City:         "Sun City",
			State:        "CA",
			Zip:          54321,
			Country:      "United States",
		},
	}
	js, err := json.Marshal(sampleResponse)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.Write(js)
}
