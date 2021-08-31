package transaction

import "models/operation"

type CreateTransactionForm struct {
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float32 `json:"amount"`
}

func (form CreateTransactionForm) ToTransaction() Transaction {
	return Transaction{
		AccountId:       form.AccountId,
		OperationTypeId: form.OperationTypeId,
		OperationType:   operation.NewOperationType(form.OperationTypeId),
		Amount:          form.Amount,
	}
}
