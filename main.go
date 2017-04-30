package main

import (
	"fmt"
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
	router.HandleFunc("/songs/{songId}", GetSong)
	//Listen and serve
	test()
	port := "8080"
	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
