package app

import (
	"log"
	"net/http"
)

func Start() {
	//define routes
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost: 8000", nil))
}


