package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type PlayListObject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type PlayListsObject []PlayListObject

func buildPlaylistsObject(rows *sql.Rows) PlayListsObject {
	var playlists PlayListsObject
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		checkErr(err)
		playlist := PlayListObject{
			Id:   id,
			Name: name,
		}
		playlists = append(playlists, playlist)
	}
	return playlists
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
func CreatePlaylist(w http.ResponseWriter, _ *http.Request) {
	//Fake playlist
	playlist := PlayListObject{
		Name: "New playlist",
	}
	playlist.Id = InsertPlaylist(playlist)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playlist)
}
func EditPlaylist(w http.ResponseWriter, r *http.Request) {

}
