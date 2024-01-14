package handler

import (
	"banking/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	custSrv services.CustomerService
}

func NewCustomerHandler(custSrv services.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSrv.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customers, err := h.custSrv.GetCustomer(id)
	if err != nil {
		switch err {
		case services.ErrCustomerNotFound:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, err)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
