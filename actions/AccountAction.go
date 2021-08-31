package actions

import (
	"models/account"
)

type IAccountRepository interface {
	Insert(account *account.Account)
	FetchByAccountId(accountId int) (bool, *account.Account)
}

type AccountAction struct {
	repository IAccountRepository
}

func NewAccountAction(repository IAccountRepository) *AccountAction {
	return &AccountAction{
		repository: repository,
	}
}

func (action *AccountAction) CreateAccount(form account.CreateAccountForm) (bool, string) {
	account := form.ToAccount()

	isValid, msg := account.IsValid()
	if !isValid {
		return false, msg
	}

	action.repository.Insert(&account)

	return true, "Account registered"
}

func (action *AccountAction) FindAccountById(accountId int) (bool, account.AccountDto) {
	exist, _account := action.repository.FetchByAccountId(accountId)
	if !exist {
		return false, account.AccountDto{}
	}

	return true, account.NewAccountDto(*_account)
}
