package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Kailash-Y/go-banking/service"
)
	

type CustomerHandlers struct {
	service service.CustomerService
}

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Kailash", City: "New York", Zipcode: "3000"},
	// 	{Name: "Yadav", City: "Bangalore", Zipcode: "560029"},
	// }
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}
}

