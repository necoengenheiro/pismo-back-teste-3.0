package repository

import "models/account"

type AccountRepository struct {
	accounts []*account.Account
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		accounts: []*account.Account{},
	}
}

func (repository *AccountRepository) Insert(account *account.Account) {
	account.Id = len(repository.accounts) + 1
	repository.accounts = append(repository.accounts, account)
}

func (repository *AccountRepository) FetchByAccountId(accountId int) (bool, *account.Account) {
	for _, account := range repository.accounts {
		if account.Id == accountId {
			return true, account
		}
	}

	return false, nil
}

func (repository *AccountRepository) ExistByAccountId(accountId int) bool {
	for _, account := range repository.accounts {
		if account.Id == accountId {
			return true
		}
	}

	return false
}
