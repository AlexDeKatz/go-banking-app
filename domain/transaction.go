package domain

import "github.com/AlexDeKatz/banking/dto"

type Transaction struct {
	TransactionID   string  `db:"transaction_id"`
	AccountID       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t *Transaction) ToTransactionResponseDTO() *dto.TransactionResponse {
	return &dto.TransactionResponse{
		TransactionId:   t.TransactionID,
		AccountId:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t *Transaction) IsWithdrawal() bool {
	return t.TransactionType == "withdrawal"
}
