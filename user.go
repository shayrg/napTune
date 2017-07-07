package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

type LoginObject struct {
	Email     string
	Password  string
	SessionId string
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
		var sessionId string
		err := rows.Scan(&email, &password, &sessionId)
		checkErr(err)
		login = LoginObject{
			Email:     email,
			Password:  password,
			SessionId: sessionId,
		}
	}
	return login
}
func Logout(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/logout.html")
}
func Login(w http.ResponseWriter, r *http.Request) {
	//Check login cookie
	storeAssets := sessions.NewCookieStore([]byte("napTune_login_session"))
	session, err := storeAssets.Get(r, "login")
	checkErr(err)
	sessionId := session.Values["sessionId"]
	userName := session.Values["userName"]
	if sessionId != nil && userName != nil {
		if userName.(string) != "" && sessionId.(string) != "" {
			//Check session and username
			loginObject := GetLogin(LoginObject{Email: userName.(string)})
			if sessionId == loginObject.SessionId {
				http.Redirect(w, r, "/", 301)
			} else {
				http.ServeFile(w, r, "./web/login.html")
			}
		} else {
			http.ServeFile(w, r, "./web/login.html")
		}
	} else {
		http.ServeFile(w, r, "./web/login.html")
	}
}
func LogoutSubmit(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("napTune_login_session"))
	session, err := store.Get(r, "login")
	checkErr(err)
	//clear database session
	if session.Values["sessionId"] != nil {
		ClearSessionId(session.Values["sessionId"].(string))
	}
	//Set cookie
	session.Values["sessionId"] = ""
	session.Values["userName"] = ""
	session.Save(r, w)
	fmt.Println("logout submited")
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
		//Set database session
		sessionId := SetSessionId(userLogin)
		//Set cookie
		var store = sessions.NewCookieStore([]byte("napTune_login_session"))
		session, err := store.Get(r, "login")
		checkErr(err)
		session.Values["sessionId"] = sessionId
		session.Values["userName"] = userLogin.Email
		session.Save(r, w)
		http.Redirect(w, r, "/", 301)
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
