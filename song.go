package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type SongStruct struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Length   string `json:"length"`
	Location string `json:"location"`
}

type SongsStruct []SongStruct

func Songs(w http.ResponseWriter, r *http.Request) {
	mysongs := GetAllSongs()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mysongs)
}
func GetSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	song := vars["songId"]
	http.ServeFile(w, r, "./music/"+song+".mp3")
}
