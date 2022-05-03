package app

import (
	"encoding/json"
	"net/http"
	"stephanusnugraha/udemy-rest-api-golang/dto"
	"stephanusnugraha/udemy-rest-api-golang/service"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (ah *AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	//decode incoming request
	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		// build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := ah.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
