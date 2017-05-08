package main

import (
	"database/sql"
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

func buildSongsObject(rows *sql.Rows) SongsObject {
	var songs SongsObject
	for rows.Next() {
		var id string
		var name string
		var artist string
		var length string
		var location string
		err := rows.Scan(&id, &name, &artist, &length, &location)
		checkErr(err)
		song := SongObject{
			Id:       id,
			Name:     name,
			Artist:   artist,
			Length:   length,
			Location: location,
		}
		songs = append(songs, song)
	}
	return songs
}
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
func UploadSong(w http.ResponseWriter, _ *http.Request) {
	//Fake song
	song := SongObject{
		Name:     "hi",
		Artist:   "Steve",
		Length:   "2 min",
		Location: "No location",
	}
	song.Id = InsertSong(song)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}
