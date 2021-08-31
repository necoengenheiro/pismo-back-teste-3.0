package account

type CreateAccountForm struct {
	DocumentNumber string `json:"document_number"`
}

func NewCreateAccountForm(documentNumber string) CreateAccountForm {
	return CreateAccountForm{
		DocumentNumber: documentNumber,
	}
}

func (form CreateAccountForm) ToAccount() Account {
	return *NewAccount(form.DocumentNumber)
}
