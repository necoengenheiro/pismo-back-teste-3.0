package repository

import (
	"models/transaction"
)

type TransactionRepository struct {
	transactions []*transaction.Transaction
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		transactions: []*transaction.Transaction{},
	}
}

func (repository *TransactionRepository) Insert(transaction *transaction.Transaction) {
	transaction.Id = len(repository.transactions) + 1
	repository.transactions = append(repository.transactions, transaction)
}
