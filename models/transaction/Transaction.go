package transaction

import (
	"models/account"
	"models/operation"
	"time"
)

type Transaction struct {
	Id              int
	Account         *account.Account
	AccountId       int
	OperationType   *operation.OperationType
	OperationTypeId int
	Amount          float32
	EventDate       time.Time
}

func (trans *Transaction) IsValid() (bool, string) {
	if !trans.OperationType.IsValid() {
		return false, "Invalid Operation"
	}

	if trans.OperationTypeId == operation.PAGAMENTO {
		if trans.Amount < 0 {
			return false, "Invalid Amount"
		}
	} else {
		if trans.Amount > 0 {
			return false, "Invalid Amount"
		}
	}

	return true, ""
}
