package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type LoginObject struct {
	Email    string
	Password string
}
type LoginResponse struct {
	LoginStatus int    `json:"loginstatus"`
	Message     string `json:"message"`
}

func buildLogin(rows *sql.Rows) LoginObject {
	var login LoginObject
	for rows.Next() {
		var email string
		var password string
		err := rows.Scan(&email, &password)
		checkErr(err)
		login = LoginObject{
			Email:    email,
			Password: password,
		}
	}
	return login
}
func Index(w http.ResponseWriter, r *http.Request) {
	//Check login cookie
	//http.ServeFile(w, r, "./web/index.html")
	//Else
	Login(w, r)
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
	var response LoginResponse
	var userLogin = LoginObject{
		Email:    r.Form["email"][0],
		Password: r.Form["password"][0],
	}
	dbLogin := GetLogin(userLogin)
	success := userExists(dbLogin)
	if success {
		success = correctPassword(dbLogin, userLogin)
		if success {
			//Success
			response = LoginResponse{
				LoginStatus: http.StatusOK,
				Message:     "Successful Login",
			}
		} else {
			//Fail
			response = LoginResponse{
				LoginStatus: http.StatusForbidden,
				Message:     "Password Incorrect",
			}
		}
	} else {
		//Fail
		response = LoginResponse{
			LoginStatus: http.StatusPreconditionFailed,
			Message:     "Email Does Not Exist In System",
		}
	}
	//Return Response
	if response.LoginStatus == http.StatusOK {
		//Set cookie
		//Index()
		http.ServeFile(w, r, "./web/index.html")
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(response.LoginStatus)
		json.NewEncoder(w).Encode(response)
	}
}
func userExists(dbLogin LoginObject) bool {
	if dbLogin.Email != "" {
		return true
	}
	return false
}
func correctPassword(dbLogin LoginObject, userLogin LoginObject) bool {
	if dbLogin.Password == userLogin.Password {
		return true
	} else {
		return false
	}
}
