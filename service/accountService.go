package service

import (
	"time"

	"github.com/AlexDeKatz/banking/domain"
	"github.com/AlexDeKatz/banking/dto"
	"github.com/AlexDeKatz/banking/errors"
)

type AccountService interface {
	CreateAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (das DefaultAccountService) CreateAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	accountData := domain.Account{
		AccountID:   "", // This will be set by the repository
		CustomerID:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
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

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
