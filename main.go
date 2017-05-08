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
	//Assets
	router.
		Methods("POST").
		Path("/login").
		Name("submitLogin").
		Handler(http.HandlerFunc(LoginSubmit))
	router.HandleFunc("/logout", LogoutSubmit)
	router.HandleFunc("/", Index)
	router.HandleFunc("/login", Login)
	router.HandleFunc("/getstyle", GetStyle)
	router.HandleFunc("/getscript", GetScript)
	//Songs
	router.
		Methods("POST").
		Path("/songs").
		Name("UploadSong").
		Handler(http.HandlerFunc(UploadSong))
	router.HandleFunc("/songs", GetSongs)
	router.HandleFunc("/songs/{songId}", GetSong)
	router.HandleFunc("/songs/{songId}/play", PlaySong)
	//Playlist
	router.
		Methods("POST").
		Path("/playlists").
		Name("CreatePlaylist").
		Handler(http.HandlerFunc(CreatePlaylist))
	router.
		Methods("PUT").
		Path("/playlists/{playlistId}").
		Name("EditPlaylist").
		Handler(http.HandlerFunc(EditPlaylist))
	router.HandleFunc("/playlists", GetPlaylists)
	router.HandleFunc("/playlists/{playlistId}", GetPlaylist)
	router.HandleFunc("/playlists/{playlistId}/songs", GetPlaylistSongs)
	//Listen and serve
	port := "8080"
	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
