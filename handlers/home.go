package handlers

import (
	"encoding/json"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"msg": "Home paga"}
	json.NewEncoder(w).Encode(res)
	return
}
