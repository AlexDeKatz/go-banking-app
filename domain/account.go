package domain

import (
	"github.com/AlexDeKatz/banking/dto"
	"github.com/AlexDeKatz/banking/errors"
)

type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(account Account) (*Account, *errors.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errors.AppError)
	FindById(accountId string) (*Account, *errors.AppError)
}

func (a *Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}

func (a *Account) ToNewAccountResponseDTO() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountID}
}
