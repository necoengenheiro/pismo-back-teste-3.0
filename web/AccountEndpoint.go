package web

import (
	"encoding/json"
	"models/account"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AccountEndpoint struct {
	router *mux.Router
	create func(account.CreateAccountForm) (bool, string)
	find   func(int) (bool, account.AccountDto)
}

func NewAccountEndpoint(router *mux.Router) *AccountEndpoint {
	endpoint := &AccountEndpoint{
		router: router,
	}

	endpoint.initRoutes()

	return endpoint
}

func (endpoint *AccountEndpoint) initRoutes() {
	endpoint.router.HandleFunc("/accounts", endpoint.createAccount).Methods("POST")
	endpoint.router.HandleFunc("/account/{accountId}", endpoint.findAccount).Methods("GET")
}

func (endpoint *AccountEndpoint) OnCreateAccount(fn func(account.CreateAccountForm) (bool, string)) {
	endpoint.create = fn
}

func (endpoint *AccountEndpoint) OnFindAccount(fn func(int) (bool, account.AccountDto)) {
	endpoint.find = fn
}

func (endpoint *AccountEndpoint) createAccount(w http.ResponseWriter, r *http.Request) {
	var accountForm account.CreateAccountForm
	json.NewDecoder(r.Body).Decode(&accountForm)

	if endpoint.create == nil {
		return
	}

	success, msg := endpoint.create(accountForm)
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

func (endpoint *AccountEndpoint) findAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountIdStr := params["accountId"]
	accountId, err := strconv.Atoi(accountIdStr)

	if endpoint.find == nil {
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	exist, account := endpoint.find(accountId)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}
