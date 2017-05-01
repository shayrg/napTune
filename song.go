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

type PlayListObject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type PlayListsObject []PlayListObject

func GetSongs(w http.ResponseWriter, _ *http.Request) {
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
func GetPlaylists(w http.ResponseWriter, _ *http.Request) {
	playlists := GetAllPlaylists()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playlists)
}
func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistId := vars["playlistId"]
	playlist := GetPlaylistById(playlistId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playlist)
}
func GetPlaylistSongs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistId := vars["playlistId"]
	songs := GetPlaylistSongsById(playlistId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}
