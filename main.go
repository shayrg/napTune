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
	router.Methods("POST").
		Path("/login").
		Name("submitLogin").
		Handler(http.HandlerFunc(LoginSubmit))
	router.HandleFunc("/logout", LogoutSubmit)
	router.HandleFunc("/", Index)
	router.HandleFunc("/login", Login)
	router.HandleFunc("/getstyle", GetStyle)
	router.HandleFunc("/getscript", GetScript)
	//Songs
	router.Methods("POST").
		Path("/songs").
		Name("UploadSong").
		Handler(http.HandlerFunc(UploadSong))
	//Edit song
	router.Methods("PUT").
		Path("/songs/{songId}").
		Name("EditSong").
		Handler(http.HandlerFunc(EditSong))
	//Delete song
	router.Methods("DELETE").
		Path("/songs/{songId}").
		Name("DelteSong").
		Handler(http.HandlerFunc(RemoveSong))
	router.HandleFunc("/songs", GetSongs)
	router.HandleFunc("/songs/{songId}", GetSong)
	router.HandleFunc("/songs/{songId}/play", PlaySong)
	//Playlist
	//All playlists
	router.Methods("POST").
		Path("/playlists").
		Name("CreatePlaylist").
		Handler(http.HandlerFunc(CreatePlaylist))
	router.HandleFunc("/playlists", GetPlaylists)
	//Playlist by id
	router.Methods("PUT").
		Path("/playlists/{playlistId}").
		Name("EditPlaylist").
		Handler(http.HandlerFunc(EditPlaylist))
	router.Methods("DELETE").
		Path("/playlists/{playlistId}").
		Name("DeletePlaylist").
		Handler(http.HandlerFunc(RemovePlaylist))
	router.HandleFunc("/playlists/{playlistId}", GetPlaylist)
	//Songs by playlist
	router.Methods("POST").
		Path("/playlists/{playlistId}/songs").
		Name("AddSongToPlaylist").
		Handler(http.HandlerFunc(AddSongToPlaylist))
	router.Methods("DELETE").
		Path("/playlists/{playlistId}/songs").
		Name("DeleteSongFromPlaylist").
		Handler(http.HandlerFunc(RemoveSongFromPlaylist))
	router.HandleFunc("/playlists/{playlistId}/songs", GetPlaylistSongs)
	//Listen and serve
	port := "8080"
	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
