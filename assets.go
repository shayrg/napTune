package main

import (
	"encoding/json"
	"net/http"
)

type LoginObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}
func GetStyle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/assets/css/style.css")
}
func GetScript(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/assets/js/javascript.js")
}
func Login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/login.html")
}
func LoginSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email string = r.Form["email"][0]
	var password string = r.Form["password"][0]
	login := LoginObject{
		Email:    email,
		Password: password,
	}
	//Check Email exists
	success := GetUserEmail(email)
	if !success {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(login)
	}
	//Check password

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(login)
}
