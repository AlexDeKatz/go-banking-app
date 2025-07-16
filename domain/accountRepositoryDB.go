package domain

import (
	"strconv"

	"github.com/AlexDeKatz/banking/errors"
	"github.com/AlexDeKatz/banking/logging"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (ard *AccountRepositoryDB) Save(account Account) (*Account, *errors.AppError) {
	insertSQL := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := ard.client.Exec(insertSQL, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logging.Error("Error while creating new account " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logging.Error("Error while getting last insert ID " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	account.AccountID = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) *AccountRepositoryDB {
	return &AccountRepositoryDB{
		client: dbClient,
	}
}
