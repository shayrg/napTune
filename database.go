package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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

//Create Song
func InsertSong(song SongObject) string {
	song.Id = getNextId("songs")
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into songs values (?, ?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(song.Id, song.Name, song.Artist, song.Length,
		song.Location)
	checkErr(err)
	db.Close()
	return song.Id
}

//Playlists table
func GetAllPlaylists() PlaylistsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	rows, err := db.Query("select * from playlists")
	checkErr(err)
	//Build Songs Object
	playlists := buildPlaylistsObject(rows)
	db.Close()
	return playlists
}
func GetPlaylistById(playlistId string) PlaylistsObject {
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

//Insert playlist
func InsertPlaylist(playlist PlaylistObject) string {
	playlist.Id = getNextId("playlists")
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into playlists values (?, ?)")
	checkErr(err)
	_, err = stmt.Exec(playlist.Id, playlist.Name)
	checkErr(err)
	db.Close()
	return playlist.Id
}
func InsertSongInPlaylist(playlistSong PlaylistSongObject) string {
	playlistSong.SongOrder = getNextPlaylistOrder(playlistSong.PlaylistId)
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into playlistSongs values (?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(playlistSong.PlaylistId, playlistSong.SongId,
		playlistSong.SongOrder)
	checkErr(err)
	db.Close()
	return playlistSong.SongOrder
}

//User Table
func GetLogin(userLogin LoginObject) LoginObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select email, password from users where email = ?")
	checkErr(err)
	rows, err := stmt.Query(userLogin.Email)
	checkErr(err)
	login := buildLogin(rows)
	db.Close()
	return login
}

//Helper functions
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getNextId(table string) string {
	queryString := "select id from " + table + " order by id desc limit 1"
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	rows, err := db.Query(queryString)
	checkErr(err)
	var id string
	for rows.Next() {
		var rowId string
		rows.Scan(&rowId)
		id = rowId
	}
	incrementString(&id)
	db.Close()
	return id
}
func getNextPlaylistOrder(playlistId string) string {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select songOrder from playlistSongs where " +
		"playlistId = ? order by songOrder desc limit 1")
	checkErr(err)
	rows, err := stmt.Query(playlistId)
	checkErr(err)
	var songOrder string
	for rows.Next() {
		var rowSongOrder string
		rows.Scan(&rowSongOrder)
		songOrder = rowSongOrder
	}
	incrementString(&songOrder)
	db.Close()
	return songOrder
}
func incrementString(string *string) {
	myInt, _ := strconv.ParseInt(*string, 0, 64)
	myInt++
	*string = fmt.Sprintf("%04d", myInt)
}
