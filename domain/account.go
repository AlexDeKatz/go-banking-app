package domain

import (
	"github.com/AlexDeKatz/banking/dto"
	"github.com/AlexDeKatz/banking/errors"
)

type Account struct {
	AccountID   string  `json:"account_id"`
	CustomerID  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type AccountRepository interface {
	Save(account Account) (*Account, *errors.AppError)
}

func (a *Account) ToNewAccountResponseDTO() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountID}
}
