package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/AlexDeKatz/banking/errors"
	"github.com/AlexDeKatz/banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		errorData := errors.NewInternalServerError("Unexpected error occurred")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorData)
		return
	}

	if customers == nil {
		errorData := errors.NewNotFoundError("Customers not found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorData)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func (ch *CustomerHandler) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := ch.service.GetCustomerById(customerId)

	if err != nil {
		errorData := errors.NewInternalServerError("Unexpected error occurred")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorData)
		return
	}

	if customer == nil {
		errorData := errors.NewNotFoundError("Customer not found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorData)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
