package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Customer details
type Customer struct {
	Name    string `json: full_name`
	City    string `json: city`
	ZipCode string `json: zip_code`
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	var customer = []Customer{
		{"Pratik", "Pune", "1234"},
		{"Pratik2", "Kutch", "2121"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

}
