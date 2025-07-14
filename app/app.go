package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlexDeKatz/banking/config"
	"github.com/AlexDeKatz/banking/domain"
	"github.com/AlexDeKatz/banking/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	config := config.GetConfig()
	if config.DatabaseURI == "" {
		log.Fatal("Database URI is not set in the environment variables")
	}
	if config.ServerPort == "" {
		log.Fatal("Server port is not set in the environment variables")
	}
	if config.ServerHost == "" {
		log.Fatal("Server host is not set in the environment variables")
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	ch := &CustomerHandler{
		// service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
		service: service.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	config := config.GetConfig()

	error := http.ListenAndServe(fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort), router)

	if error != nil {
		log.Fatal(error)
	}
}
