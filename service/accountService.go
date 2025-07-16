package service

import (
	"time"

	"github.com/AlexDeKatz/banking/domain"
	"github.com/AlexDeKatz/banking/dto"
	"github.com/AlexDeKatz/banking/errors"
)

type AccountService interface {
	CreateAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

const dbTSLayout = "2006-01-02 15:04:05"

func (das DefaultAccountService) CreateAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	accountData := domain.Account{
		AccountID:   "", // This will be set by the repository
		CustomerID:  request.CustomerId,
		OpeningDate: time.Now().Format(dbTSLayout),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}
	newAccount, err := das.repo.Save(accountData)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDTO()

	return response, nil
}

func (das DefaultAccountService) MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if request.IsTransactionTypeWithdrawal() {
		account, err := das.repo.FindById(request.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(request.Amount) {
			return nil, errors.NewValidationError("Insufficient balance in the account")
		}
	}

	transactionData := domain.Transaction{
		AccountID:       request.AccountId,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}

	transaction, err := das.repo.SaveTransaction(transactionData)
	if err != nil {
		return nil, err
	}

	response := transaction.ToTransactionResponseDTO()
	return response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
