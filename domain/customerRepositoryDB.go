package domain

import (
	"database/sql"
	"time"

	"github.com/AlexDeKatz/banking/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (crd *CustomerRepositoryDB) FindAll(status string) ([]Customer, error) {
	customers := make([]Customer, 0)
	var err error
	if status == "" {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		err = crd.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?"
		err = crd.client.Select(&customers, findAllSQL, status)
	}

	if err != nil {
		logging.Error("Error while querying customers " + err.Error())
		return nil, err
	}
	return customers, nil
}

func (crd *CustomerRepositoryDB) FindById(id string) (*Customer, error) {
	findByIdSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	var customer Customer
	err := crd.client.Get(&customer, findByIdSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No customer found
		}
		logging.Error("Error while scanning customer by ID " + err.Error())
		return nil, err
	}
	return &customer, nil
}

func NewCustomerRepositoryDB() *CustomerRepositoryDB {
	db, err := sqlx.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &CustomerRepositoryDB{client: db}
}
