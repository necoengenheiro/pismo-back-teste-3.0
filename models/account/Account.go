package account

type Account struct {
	Id             int
	DocumentNumber string
}

func NewAccount(documentNumber string) *Account {
	return &Account{
		Id:             0,
		DocumentNumber: documentNumber,
	}
}

func (account *Account) IsValid() (bool, string) {
	if len(account.DocumentNumber) == 0 {
		return false, "Informe o número do documento"
	}

	return true, ""
}
