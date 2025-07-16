package app

import (
	"encoding/json"
	"net/http"

	"github.com/AlexDeKatz/banking/dto"
	"github.com/AlexDeKatz/banking/logging"
	"github.com/AlexDeKatz/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest

	queryParams := mux.Vars(r)
	request.CustomerId = queryParams["customer_id"]

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		account, err := ah.service.CreateAccount(request)
		if err != nil {
			w.WriteHeader(err.Code)
			json.NewEncoder(w).Encode(err.Message)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)
	}
}

func (ah *AccountHandler) makeTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.TransactionRequest

	queryParams := mux.Vars(r)
	request.CustomerId = queryParams["customer_id"]
	request.AccountId = queryParams["account_id"]

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logging.Error("Error while decoding: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		transaction, err := ah.service.MakeTransaction(request)
		if err != nil {
			logging.Error("Error while making transaction: " + err.Message)
			w.WriteHeader(err.Code)
			json.NewEncoder(w).Encode(err.Message)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(transaction)
	}
}
