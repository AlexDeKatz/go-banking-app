package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AlexDeKatz/banking/config"
	"github.com/AlexDeKatz/banking/domain"
	"github.com/AlexDeKatz/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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

	dbClient := getDBClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)

	ch := &CustomerHandler{
		// service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
		service: service.NewCustomerService(customerRepositoryDB),
	}

	ah := &AccountHandler{
		service: service.NewAccountService(accountRepositoryDB),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.createAccount).Methods(http.MethodPost)

	config := config.GetConfig()

	error := http.ListenAndServe(fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort), router)

	if error != nil {
		log.Fatal(error)
	}
}

func getDBClient() *sqlx.DB {
	config := config.GetConfig()
	db, err := sqlx.Open("mysql", config.DatabaseURI)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
