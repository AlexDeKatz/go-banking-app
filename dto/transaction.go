package dto

import "github.com/AlexDeKatz/banking/errors"

type TransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (tr *TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return tr.TransactionType == "withdrawal"
}

func (tr *TransactionRequest) IsTransactionTypeDeposit() bool {
	return tr.TransactionType == "deposit"
}

func (tr *TransactionRequest) Validate() *errors.AppError {
	if !tr.IsTransactionTypeDeposit() && !tr.IsTransactionTypeWithdrawal() {
		return errors.NewValidationError("Transaction type must be either 'withdrawal' or 'deposit'")
	}
	if tr.Amount <= 0 {
		return errors.NewValidationError("Amount must be greater than zero")
	}
	return nil
}
