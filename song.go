package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type SongStruct struct {
	Name     string `json:"name"`
	Length   string `json:"length"`
	Location string `json:"location"`
	Order    string `json:"order"`
}

type SongsStruct []SongStruct

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
func GetSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	song := vars["songId"]
	http.ServeFile(w, r, "./music/"+song+".mp3")
}
