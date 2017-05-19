package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type PlaylistSongObject struct {
	PlaylistId string `json:"playlist"`
	SongId     string `json:"song"`
	SongOrder  string `json:"order"`
}
type PlaylistSongsObject []PlaylistSongObject

func buildPlaylistSongsObject(rows *sql.Rows) PlaylistSongsObject {
	var playlistSongs PlaylistSongsObject
	for rows.Next() {
		var playlistId string
		var songId string
		var songOrder string
		err := rows.Scan(&playlistId, &songId, &songOrder)
		checkErr(err)
		playlistSong := PlaylistSongObject{
			PlaylistId: playlistId,
			SongId:     songId,
			SongOrder:  songOrder,
		}
		playlistSongs = append(playlistSongs, playlistSong)
	}
	return playlistSongs
}
func GetPlaylistSongs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistId := vars["playlistId"]
	songs := GetPlaylistSongsByPlaylistId(playlistId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}
func GetPlaylistSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistSong := PlaylistSongObject{
		SongId:     vars["songId"],
		PlaylistId: vars["playlistId"],
	}
	songs := GetPlaylistSongsById(playlistSong)
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
	playlistSongs := InsertSongInPlaylist(playlistSong)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playlistSongs)
}
func EditSongInPlaylist(w http.ResponseWriter, r *http.Request) {
	//Fake playlistSongObject
	playlistSong := PlaylistSongObject{
		SongOrder: "9999",
	}
	vars := mux.Vars(r)
	playlistSong.PlaylistId = vars["playlistId"]
	playlistSong.SongId = vars["songId"]
	rowsAffected := UpdateSongInPlaylist(playlistSong)
	if rowsAffected >= 1 {
		GetPlaylistSong(w, r)
	}
}
func RemoveSongFromPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistSong := PlaylistSongObject{
		SongId:     vars["songId"],
		PlaylistId: vars["playlistId"],
	}
	rowsAffected := DeleteSongInPlaylist(playlistSong)
	//List Playlist songs
	if rowsAffected >= 1 {
		GetPlaylistSongs(w, r)
	}
}
