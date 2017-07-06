package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
)

const dbString string = "root:@tcp(localhost:3306)" +
	"/napTune?charset=utf8"

/**
 *Songs table
 */
//Get All
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

//Get
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

//Insert
func InsertSong(song SongObject) SongsObject {
	song.Id = getNextId("songs")
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into songs values (?, ?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(song.Id, song.Name, song.Artist, song.Length,
		song.Location)
	checkErr(err)
	db.Close()
	var songs SongsObject
	songs = append(songs, song)
	return songs
}

//Update
func UpdateSong(song SongObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("update songs set name = ?, artist = ?, " +
		"length = ?, location = ? where id = ? ")
	checkErr(err)
	result, err := stmt.Exec(song.Name, song.Artist, song.Length,
		song.Location, song.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

//Delete
func DeleteSong(song SongObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from songs where " +
		"id = ?")
	checkErr(err)
	result, err := stmt.Exec(song.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

/**
 *Playlists table
 */
//Get All
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

//Get
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

//Insert
func InsertPlaylist(playlist PlaylistObject) PlaylistsObject {
	playlist.Id = getNextId("playlists")
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into playlists values (?, ?)")
	checkErr(err)
	_, err = stmt.Exec(playlist.Id, playlist.Name)
	checkErr(err)
	db.Close()
	var playlists PlaylistsObject
	playlists = append(playlists, playlist)
	return playlists
}

//Update
func UpdatePlaylist(playlist PlaylistObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("update playlists set name = ? where id = ?")
	checkErr(err)
	result, err := stmt.Exec(playlist.Name, playlist.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

//Delete
func DeletePlaylist(playlist PlaylistObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from playlists where id = ?")
	checkErr(err)
	result, err := stmt.Exec(playlist.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

/**
 *PlaylistSongs Table
 */
//Get all in playlist
func GetPlaylistSongsByPlaylistId(playlistId string) SongsObject {
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

//Get
func GetPlaylistSongsById(playlistSong PlaylistSongObject) PlaylistSongsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select * from playlistSongs where songId = ?" +
		" and playlistId = ?")
	checkErr(err)
	rows, err := stmt.Query(playlistSong.SongId, playlistSong.PlaylistId)
	checkErr(err)
	db.Close()
	playlistSongs := buildPlaylistSongsObject(rows)
	return playlistSongs
}

//Insert
func InsertSongInPlaylist(playlistSong PlaylistSongObject) PlaylistSongsObject {
	playlistSong.SongOrder = getNextPlaylistOrder(playlistSong.PlaylistId)
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into playlistSongs values (?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(playlistSong.PlaylistId, playlistSong.SongId,
		playlistSong.SongOrder)
	checkErr(err)
	db.Close()
	var playlistSongs PlaylistSongsObject
	playlistSongs = append(playlistSongs, playlistSong)
	return playlistSongs
}

//Update
func UpdateSongInPlaylist(playlistSong PlaylistSongObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("update playlistSongs set songOrder = ? where" +
		" songId = ? and playlistId = ?")
	checkErr(err)
	result, err := stmt.Exec(playlistSong.SongOrder, playlistSong.SongId,
		playlistSong.PlaylistId)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

//Delete
func DeleteSongInPlaylist(playlistSong PlaylistSongObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from playlistSongs where " +
		"playlistId = ? and songId = ?")
	checkErr(err)
	result, err := stmt.Exec(playlistSong.PlaylistId, playlistSong.SongId)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}
func DeleteAllSongsInPlaylist(playlist PlaylistObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from playlistSongs where " +
		"playlistId = ?")
	checkErr(err)
	result, err := stmt.Exec(playlist.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}
func DeleteSongFromAllPlaylists(song SongObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from playlistSongs where " +
		"songId = ?")
	checkErr(err)
	result, err := stmt.Exec(song.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

/**
 *User Table
 */
//Get
func GetLogin(userLogin LoginObject) LoginObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select email, password, sessionId from users" +
		" where email = ?")
	checkErr(err)
	rows, err := stmt.Query(userLogin.Email)
	checkErr(err)
	login := buildLogin(rows)
	db.Close()
	return login
}
func SetSessionId(login LoginObject) string {
	string :=
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	stringLength := 32
	sessionId := make([]byte, stringLength)
	for i := range sessionId {
		sessionId[i] = string[rand.Intn(len(string))]
	}
	sessionIdString := fmt.Sprintf("%x", sessionId)
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("update users set sessionId = ? where " +
		"email = ? and password = ?")
	checkErr(err)
	_, err = stmt.Exec(sessionIdString, login.Email, login.Password)
	checkErr(err)
	db.Close()
	return sessionIdString
}
func ClearSessionId(sessionId string) {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("update users set sessionId = '' where " +
		"sessionId = ?")
	checkErr(err)
	_, err = stmt.Exec(sessionId)
	checkErr(err)
	db.Close()
}

/**
 *Helper functions
 */
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
