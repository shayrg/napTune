package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
	//"path/filepath"
	//"os"
)

func main() {
	//Make router
	router := mux.NewRouter().StrictSlash(true)
	//Routs
	router.HandleFunc("/", Index)
	router.HandleFunc("/songs", Songs)
	router.HandleFunc("/songs/{songId}", GetSongById)
	router.HandleFunc("/getpic", GetPic)
	router.HandleFunc("/getsong/{song}", GetSong)
	router.HandleFunc("/getstyle", GetStyle)
	router.HandleFunc("/getscript", GetScript)
	//Listen and serve
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}
func Songs(w http.ResponseWriter, r *http.Request) {
	mysongs := SongsStruct{
		SongStruct{Name: "Help", Length: "3 Min", Order: "2",
			Location: "getsong/song2"},
		SongStruct{Name: "Steve", Length: "2 Min", Order: "2",
			Location: "getsong/song1"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mysongs)
}
func GetSongById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songId := vars["songId"]
	mystring := "Your song id is " + songId
	fmt.Fprintln(w, mystring)
}
func GetPic(w http.ResponseWriter, r *http.Request) {
	//http.FileServer(http.Dir("./music/pic.jpg"))
	http.ServeFile(w, r, "./music/pic.jpg")
}
func GetSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	song := vars["song"]
	http.ServeFile(w, r, "./music/"+song+".mp3")
}
func GetStyle(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./web/assets/css/style.css")
}
func GetScript(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./web/assets/js/javascript.js")
}
