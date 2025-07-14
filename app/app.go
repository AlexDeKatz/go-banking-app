package app

import (
	"log"
	"net/http"

	"github.com/AlexDeKatz/banking/domain"
	"github.com/AlexDeKatz/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	ch := &CustomerHandler{
		// service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
		service: service.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	error := http.ListenAndServe(":8080", router)

	if error != nil {
		log.Fatal(error)
	}
}
