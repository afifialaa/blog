package handlers

import (
	"encoding/json"
	"net/http"
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"msg": "Home page"}
	json.NewEncoder(w).Encode(res)
	return
}
func ReadBlog(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"msg": "Home page"}
	json.NewEncoder(w).Encode(res)
	return
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"msg": "Home page"}
	json.NewEncoder(w).Encode(res)
	return
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"msg": "Home page"}
	json.NewEncoder(w).Encode(res)
	return
}
