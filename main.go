package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", greeting)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
