package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const dbString string = "root:@tcp(localhost:3306)" +
	"/napTune?charset=utf8"

//Songs table
func GetAllSongs() SongsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	rows, err := db.Query("select * from songs")
	checkErr(err)
	//Build Songs Object
	songs := buildSongsObject(rows)
	db.Close()
	return songs
}
func GetSongById(songId string) SongsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select * from songs where id = ?")
	checkErr(err)
	rows, err := stmt.Query(songId)
	checkErr(err)
	//Build Songs Object
	songs := buildSongsObject(rows)
	db.Close()
	return songs
}

//Playlists table
func GetAllPlaylists() PlayListsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	rows, err := db.Query("select * from playlists")
	checkErr(err)
	//Build Songs Object
	playlists := buildPlaylistsObject(rows)
	db.Close()
	return playlists
}
func GetPlaylistById(playlistId string) PlayListsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select * from playlists where id = ?")
	checkErr(err)
	rows, err := stmt.Query(playlistId)
	checkErr(err)
	//Build Songs Object
	playlists := buildPlaylistsObject(rows)
	db.Close()
	return playlists
}
func GetPlaylistSongsById(playlistId string) SongsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select songs.id, songs.name, artist, length, location from songs join playlistSongs on songs.id = songId join playlists on playlists.id = playlistId where playlistId = ?")
	checkErr(err)
	rows, err := stmt.Query(playlistId)
	checkErr(err)
	//Build Songs Object
	songs := buildSongsObject(rows)
	db.Close()
	return songs
}

//Helper functions
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
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
