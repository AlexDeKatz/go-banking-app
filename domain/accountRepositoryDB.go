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

func (ard *AccountRepositoryDB) FindById(accountId string) (*Account, *errors.AppError) {
	sqlGetAccount := "SELECT  account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	if err := ard.client.Get(&account, sqlGetAccount, accountId); err != nil {
		logging.Error("Error while fetching account information: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func (ard *AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errors.AppError) {
	tx, err := ard.client.Begin()
	if err != nil {
		logging.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected Database error")
	}
	insertTrxSQL := "INSERT into transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	result, _ := tx.Exec(insertTrxSQL, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? WHERE account_id = ?`, t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? WHERE account_id = ?`, t.Amount, t.AccountID)
	}

	if err != nil {
		tx.Rollback()
		logging.Error("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logging.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logging.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	account, appError := ard.FindById(t.AccountID)
	if appError != nil {
		return nil, appError
	}
	t.TransactionID = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) *AccountRepositoryDB {
	return &AccountRepositoryDB{
		client: dbClient,
	}
}
