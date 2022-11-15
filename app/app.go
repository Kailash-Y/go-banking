package app

import (
	"log"
	"net/http"

	"github.com/Kailash-Y/go-banking/domain"
	"github.com/Kailash-Y/go-banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	//define routes
	router := mux.NewRouter()
    //wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost: 8000", router))
}
