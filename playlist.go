package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//Types
type PlaylistObject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type PlaylistsObject []PlaylistObject

//Functions
func buildPlaylistsObject(rows *sql.Rows) PlaylistsObject {
	var playlists PlaylistsObject
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		checkErr(err)
		playlist := PlaylistObject{
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
func CreatePlaylist(w http.ResponseWriter, _ *http.Request) {
	//Fake playlist
	playlist := PlaylistObject{
		Name: "New playlist",
	}
	playlist.Id = InsertPlaylist(playlist)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playlist)
}

func EditPlaylist(w http.ResponseWriter, r *http.Request) {

}
func RemovePlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlist := PlaylistObject{
		Id: vars["playlistId"],
	}
	//Delete all songs from playlist
	DeleteAllSongsInPlaylist(playlist)
	//Delete playlist
	DeletePlaylist(playlist)
	//Show all playlists
	GetPlaylists(w, r)
}
