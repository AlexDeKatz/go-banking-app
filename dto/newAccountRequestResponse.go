package dto

import (
	"strings"

	"github.com/AlexDeKatz/banking/errors"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id" xml:"customer_id"`
	AccountType string  `json:"account_type" xml:"account_type"`
	Amount      float64 `json:"amount" xml:"amount"`
}

type NewAccountResponse struct {
	AccountId string `json:"account_id" xml:"account_id"`
}

func (nar NewAccountRequest) Validate() *errors.AppError {
	if nar.Amount < 5000 {
		return errors.NewValidationError("Minimum amount for account creation is 5000")
	}
	if strings.ToLower(nar.AccountType) != "saving" && strings.ToLower(nar.AccountType) != "checking" {
		return errors.NewValidationError("Account type must be either 'saving' or 'checking'")
	}
	return nil
}
