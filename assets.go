package main

import (
	"fmt"
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
	fmt.Println("username:", r.Form["email"])
	fmt.Println("password:", r.Form["password"])
}
