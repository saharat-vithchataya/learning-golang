package services

import (
	"banking/logs"
	"banking/repository"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (srv accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	// Validate
	newAcc, err := srv.accRepo.Create(repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      true,
	})
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}, nil
}

func (srv accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := srv.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return responses, nil
}
