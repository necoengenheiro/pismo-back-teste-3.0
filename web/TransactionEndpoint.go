package web

import (
	"encoding/json"
	"models/transaction"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionEndpoint struct {
	router *mux.Router
	create func(transaction.CreateTransactionForm) (bool, string)
}

func NewTransactionEndpoint(router *mux.Router) *TransactionEndpoint {
	endpoint := &TransactionEndpoint{
		router: router,
	}

	endpoint.initRoutes()

	return endpoint
}

func (endpoint *TransactionEndpoint) initRoutes() {
	endpoint.router.HandleFunc("/transactions", endpoint.createTransaction).Methods("POST")
}

func (endpoint *TransactionEndpoint) OnCreateTransaction(fn func(transaction.CreateTransactionForm) (bool, string)) {
	endpoint.create = fn
}

func (endpoint *TransactionEndpoint) createTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionForm transaction.CreateTransactionForm
	json.NewDecoder(r.Body).Decode(&transactionForm)

	if endpoint.create == nil {
		return
	}

	success, msg := endpoint.create(transactionForm)
	if !success {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Protocol{
			Status:  http.StatusBadRequest,
			Message: msg,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
