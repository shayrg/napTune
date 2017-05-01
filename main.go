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
	router.HandleFunc("/login", Login)
	router.HandleFunc("/getstyle", GetStyle)
	router.HandleFunc("/getscript", GetScript)
	//Songs
	router.HandleFunc("/songs", GetSongs)
	router.HandleFunc("/songs/{songId}", GetSong)
	router.HandleFunc("/songs/{songId}/play", PlaySong)
	router.HandleFunc("/songs/upload", UploadSong)
	//Playlist
	router.HandleFunc("/playlists", GetPlaylists)
	router.HandleFunc("/playlists/{playlistId}", GetPlaylist)
	router.HandleFunc("/playlists/{playlistId}/songs", GetPlaylistSongs)
	router.HandleFunc("/playlists/create", CreatePlaylist)
	router.HandleFunc("/playlists/edit", EditPlaylist)
	//Listen and serve
	port := "8080"
	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
