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
	/**
	 *Users
	 */
	//Login
	router.Methods("POST").
		Path("/login").
		Name("submitLogin").
		Handler(http.HandlerFunc(LoginSubmit))
	//Logout
	router.HandleFunc("/logout", LogoutSubmit)
	/**
	 *Songs
	 */
	//Create song
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
	//Get all Songs
	router.HandleFunc("/songs", GetSongs)
	//Get song by Id
	router.HandleFunc("/songs/{songId}", GetSong)
	//Play song by Id
	router.HandleFunc("/songs/{songId}/play", PlaySong)
	/**
	 *Playlist
	 */
	//Create playlist
	router.Methods("POST").
		Path("/playlists").
		Name("CreatePlaylist").
		Handler(http.HandlerFunc(CreatePlaylist))
	//Edit playlist
	router.Methods("PUT").
		Path("/playlists/{playlistId}").
		Name("EditPlaylist").
		Handler(http.HandlerFunc(EditPlaylist))
	//Delete playlist
	router.Methods("DELETE").
		Path("/playlists/{playlistId}").
		Name("DeletePlaylist").
		Handler(http.HandlerFunc(RemovePlaylist))
	//Get all playlists
	router.HandleFunc("/playlists", GetPlaylists)
	//Get playlist by id
	router.HandleFunc("/playlists/{playlistId}", GetPlaylist)
	/**
	 *PlaylistSongs
	 */
	//Add song to playlist
	router.Methods("POST").
		Path("/playlists/{playlistId}/songs").
		Name("AddSongToPlaylist").
		Handler(http.HandlerFunc(AddSongToPlaylist))
	//Edit song in playlist
	router.Methods("PUT").
		Path("/playlists/{playlistId}/songs/{songId}").
		Name("EditSongInPlaylist").
		Handler(http.HandlerFunc(EditSongInPlaylist))
	//Remove song from playlist
	router.Methods("DELETE").
		Path("/playlists/{playlistId}/songs/{songId}").
		Name("DeleteSongFromPlaylist").
		Handler(http.HandlerFunc(RemoveSongFromPlaylist))
	//Get all songs in playlist
	router.HandleFunc("/playlists/{playlistId}/songs", GetPlaylistSongs)
	//Get playlistSong by id
	router.HandleFunc("/playlists/{playlistId}/songs/{songId}",
		GetPlaylistSong)
	/**
	 *Assets
	 */
	//Serve Index
	router.HandleFunc("/", Index)
	//Serve Login
	router.HandleFunc("/login", Login)
	//Serve Style
	router.HandleFunc("/getstyle", GetStyle)
	//Serve Script
	router.HandleFunc("/getscript", GetScript)
	//Listen and serve
	port := "8080"
	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
