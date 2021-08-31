package actions

import (
	"models/transaction"
	"time"
)

type ITransactionRepository interface {
	Insert(*transaction.Transaction)
}

type TransactionAction struct {
	repository ITransactionRepository
}

func NewTransactionAction(repository ITransactionRepository) *TransactionAction {
	return &TransactionAction{
		repository: repository,
	}
}

func (action *TransactionAction) CreateTransaction(transactionFrom transaction.CreateTransactionForm) (bool, string) {
	transaction := transactionFrom.ToTransaction()

	isValidTransaction, msg := transaction.IsValid()
	if !isValidTransaction {
		return false, msg
	}

	transaction.EventDate = time.Now()
	action.repository.Insert(&transaction)

	return true, "transaction registered"
}
