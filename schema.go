package main

type SampleResponse struct {
	Hello   string      `json:"name_of_person"`
	Count   int         `json:"count"`
	Address FullAddress `json:"address"`
}

type FullAddress struct {
	StreetNumber uint64
	Street       string
	City         string
	State        string
	Zip          uint64
	Country      string
}
