package main

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}
func GetStyle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/assets/css/style.css")
}
func GetScript(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/assets/js/javascript.js")
}
