package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type SongObject struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Length   string `json:"length"`
	Location string `json:"location"`
}

type SongsObject []SongObject

func GetSongs(w http.ResponseWriter, r *http.Request) {
	songs := GetAllSongs()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}
func GetSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songId := vars["songId"]
	song := GetSongById(songId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}
func PlaySong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songId := vars["songId"]
	http.ServeFile(w, r, "./web/assets/media/music/"+songId+".mp3")
}
