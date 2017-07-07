package main

import (
	"github.com/gorilla/sessions"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//Check login cookie
	storeAssets := sessions.NewCookieStore([]byte("napTune_login_session"))
	session, err := storeAssets.Get(r, "login")
	checkErr(err)
	sessionId := session.Values["sessionId"]
	userName := session.Values["userName"]
	/*fmt.Println("SesionId")
	fmt.Println(sessionId)
	fmt.Println("username")
	fmt.Println(userName)*/
	if sessionId != nil && userName != nil {
		if userName.(string) != "" && sessionId.(string) != "" {
			//Check session and username
			loginObject := GetLogin(LoginObject{Email: userName.(string)})
			if sessionId == loginObject.SessionId {
				http.ServeFile(w, r, "./web/index.html")
			} else {
				//Login(w, r)
				http.Redirect(w, r, "/login", 301)
			}
		} else {
			http.Redirect(w, r, "/login", 301)
		}
	} else {
		http.Redirect(w, r, "/login", 301)
	}
}
func GetStyle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/assets/css/style.css")
}
func GetScript(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/assets/js/javascript.js")
}
