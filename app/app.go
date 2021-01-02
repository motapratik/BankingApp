package app

import (
	"log"
	"net/http"
)

// Start func
func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greeting)
	mux.HandleFunc("/customer", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
