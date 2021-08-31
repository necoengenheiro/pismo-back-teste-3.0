package main

import (
	"actions"
	"log"
	"models/account"
	"models/transaction"
	"net/http"
	"repository"
	"web"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	accountRepository := repository.NewAccountRepository()
	accountEndpoint := web.NewAccountEndpoint(router)
	accountAction := actions.NewAccountAction(accountRepository)

	transactionRepository := repository.NewTransactionRepository()
	transactionEndpoint := web.NewTransactionEndpoint(router)
	transactionAction := actions.NewTransactionAction(transactionRepository)

	accountEndpoint.OnCreateAccount(func(form account.CreateAccountForm) (bool, string) {
		return accountAction.CreateAccount(form)
	})

	accountEndpoint.OnFindAccount(func(accountId int) (bool, account.AccountDto) {
		return accountAction.FindAccountById(accountId)
	})

	transactionEndpoint.OnCreateTransaction(func(transactionFrom transaction.CreateTransactionForm) (bool, string) {
		return transactionAction.CreateTransaction(transactionFrom)
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}
