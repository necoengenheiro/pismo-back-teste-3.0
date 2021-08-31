package account

type AccountDto struct {
	Id             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func NewAccountDto(account Account) AccountDto {
	return AccountDto{
		Id:             account.Id,
		DocumentNumber: account.DocumentNumber,
	}
}

func (dto AccountDto) ToAccount() Account {
	return Account{
		Id:             dto.Id,
		DocumentNumber: dto.DocumentNumber,
	}
}
