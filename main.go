package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//Make router
	router := mux.NewRouter().StrictSlash(true)
	//Routs
	//Assets
	router.HandleFunc("/", Index)
	router.HandleFunc("/getstyle", GetStyle)
	router.HandleFunc("/getscript", GetScript)
	//Songs
	router.HandleFunc("/songs", Songs)
	router.HandleFunc("/songs/{songId}", GetSongById)
	router.HandleFunc("/getsong/{song}", GetSong)
	//Listen and serve
	log.Fatal(http.ListenAndServe(":8080", router))
}
