package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type PlaylistSongObject struct {
	PlaylistId string
	SongId     string
	SongOrder  string
}

func GetPlaylistSongs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistId := vars["playlistId"]
	songs := GetPlaylistSongsById(playlistId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}
func AddSongToPlaylist(w http.ResponseWriter, r *http.Request) {
	//Fake playlistSongObject
	playlistSong := PlaylistSongObject{
		SongId: "0003",
	}
	vars := mux.Vars(r)
	playlistSong.PlaylistId = vars["playlistId"]
	playlistSong.SongOrder = InsertSongInPlaylist(playlistSong)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playlistSong)
}
func RemoveSongFromPlaylist(w http.ResponseWriter, r *http.Request) {
	//Fake playlistSongObject
	playlistSong := PlaylistSongObject{
		SongId: "0002",
	}
	vars := mux.Vars(r)
	playlistSong.PlaylistId = vars["playlistId"]
	DeleteSongInPlaylist(playlistSong)
	//List Playlist songs
	GetPlaylistSongs(w, r)
}
