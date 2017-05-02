package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func userExists(rows sql.Rows) bool {
	if rows.Next() {
		return true
	}
	return false
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
	var rows sql.Rows = GetUserEmail(email)
	success := userExists(rows)
	if success {
		success = correctPassword(email, password)
		if success {

		}
		//Fail
	}
	//Fail

	/*login := LoginObject{
		Email:    email,
		Password: password,
	}
	//Check Email exists
	success := GetUserEmail(email)
	if !success {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(login)
	} else {
		//Check password
		success = GetUserInfo(login)
		if !success {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(login)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(login)
		}
	}*/
}
